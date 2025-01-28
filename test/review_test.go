package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tbourn/debugzen-backend/handlers/review"
	"github.com/tbourn/debugzen-backend/internal/config"
	"github.com/tbourn/debugzen-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
)

func TestReviewHandler(t *testing.T) {
	config.LoadEnv()

	apiKey := config.GetEnv("OPENAI_API_KEY", "")
	if apiKey == "" {
		t.Fatal("OPENAI_API_KEY is not set in environment variables")
	}

	openAIClient := openai.NewClient(apiKey)
	reviewService := services.NewReviewService(openAIClient)

	router := gin.Default()
	reviewHandler := review.NewReviewHandler(reviewService)
	router.POST("/review", reviewHandler.Review)

	t.Run("Successful Code Review", func(t *testing.T) {
		body := `{"code": "def hello_world(): print('Hello, World!')"}`
		req, _ := http.NewRequest("POST", "/review", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.NotEmpty(t, response["feedback"])
		feedback := response["feedback"].([]interface{})

		t.Logf("Feedback: %v", feedback)
	})
}