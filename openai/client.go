package openai

import (
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

var client *openai.Client

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set in environment variables")
	}
	client = openai.NewClient(apiKey)
}

func GetOpenAIClient() *openai.Client {
	return client
}