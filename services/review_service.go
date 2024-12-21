package services

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type ReviewService struct {
	Client *openai.Client
}

type Feedback struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewReviewService(client *openai.Client) *ReviewService {
	return &ReviewService{
		Client: client,
	}
}

func (r *ReviewService) GetCodeReviewFeedback(code string) ([]Feedback, error) {
	prompt := `Review the following code for errors, improvements, and best practices:

` + code + `

Provide a detailed feedback list in the following format:
1. Title of feedback
   - Description explaining the feedback in detail.
2. Title of feedback
   - Description explaining the feedback in detail.
...

Each suggestion should be structured with a numbered title followed by a description. Avoid unnecessary text or unrelated information.`

	if r.Client == nil {
		return nil, fmt.Errorf("OpenAI client is not initialized")
	}

	req := openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are a code review assistant. Your task is to provide structured and constructive feedback.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		MaxTokens:   500,
		Temperature: 0.7,
	}

	resp, err := r.Client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Printf("Error calling OpenAI API: %v", err)
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no feedback received from OpenAI")
	}

	rawFeedback := strings.Split(resp.Choices[0].Message.Content, "\n")
	feedback := []Feedback{}
	var current Feedback

	for _, line := range rawFeedback {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "1.") || strings.HasPrefix(trimmed, "2.") ||
			strings.HasPrefix(trimmed, "3.") || strings.HasPrefix(trimmed, "4.") {
			if current.Title != "" || current.Description != "" {
				feedback = append(feedback, current)
			}
			current = Feedback{Title: strings.TrimSpace(trimmed), Description: ""}
		} else if strings.HasPrefix(trimmed, "- ") {
			current.Description += strings.TrimPrefix(trimmed, "- ") + "\n"
		}
	}

	if current.Title != "" || current.Description != "" {
		feedback = append(feedback, current)
	}

	return feedback, nil
}