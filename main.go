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
		AllowOrigins:     []string{config.GetEnv("BASE_URL", "http://localhost:5173"), "https://tbourn.github.io"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/review", review.NewReviewHandler(reviewService))

	port := config.GetEnv("PORT", "8080")
	log.Printf("Starting server on port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}