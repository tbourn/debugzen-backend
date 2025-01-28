package main

import (
	"log"

	"github.com/tbourn/debugzen-backend/handlers/review"
	"github.com/tbourn/debugzen-backend/internal/config"
	"github.com/tbourn/debugzen-backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tbourn/debugzen-backend/docs"
)

// @title DebugZen API
// @version 1.0.1
// @description API for DebugZen Backend.
// @host localhost:8080
// @BasePath /
func main() {
	config.LoadEnv()

	openAIClient := openai.NewClient(config.GetEnv("OPENAI_API_KEY", ""))
	reviewService := services.NewReviewService(openAIClient)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173", "https://tbourn.github.io"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// @Summary Home route
	// @Description Welcome message for DebugZen API
	// @Tags root
	// @Produce json
	// @Success 200 {object} map[string]string
	// @Router / [get]
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to DebugZen Backend API. Use /review to submit your code. Example: curl -X POST https://debugzen-backend.onrender.com/review -H \"Content-Type: application/json\" -d '{\"code\": \"def hello_world():\\n    print(\\\"Hello, World!\\\")\"}'",
		})
	})

	// @Summary Submit code for review
	// @Description Sends code to OpenAI for review and feedback
	// @Tags review
	// @Accept  json
	// @Produce  json
	// @Param request body review.ReviewRequest true "Code to analyze"
	// @Success 200 {object} review.ReviewResponse
	// @Failure 400 {object} review.ErrorResponse
	// @Router /review [post]
	reviewHandler := review.NewReviewHandler(reviewService)
	r.POST("/review", reviewHandler.Review)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := config.GetEnv("PORT", "8080")
	log.Printf("Starting server on port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}