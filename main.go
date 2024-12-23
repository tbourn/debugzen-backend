package main

import (
	"log"

	"github.com/tbourn/debugzen-backend/handlers/review"
	"github.com/tbourn/debugzen-backend/internal/config"
	"github.com/tbourn/debugzen-backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to DebugZen Backend API. Use /review to submit your code. Example: curl -X POST https://debugzen-backend.onrender.com/review -H \"Content-Type: application/json\" -d '{\"code\": \"def hello_world():\\n    print(\\\"Hello, World!\\\")\"}'",
		})
	})

	r.POST("/review", review.NewReviewHandler(reviewService))

	port := config.GetEnv("PORT", "8080")
	log.Printf("Starting server on port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}