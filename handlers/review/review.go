package review

import (
	"debugzen/internal/utils"
	"debugzen/services"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ReviewRequest struct {
	Code string `json:"code" binding:"required"`
}

const (
	ErrInvalidPayload   = "Invalid request payload"
	ErrEmptyCodeInput   = "Code input is required"
	ErrProcessingReview = "Failed to process the code review"
)

func NewReviewHandler(reviewService *services.ReviewService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ReviewRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, ErrInvalidPayload)
			return
		}

		code := strings.TrimSpace(req.Code)
		if code == "" {
			utils.RespondWithError(c, http.StatusBadRequest, ErrEmptyCodeInput)
			return
		}

		feedback, err := reviewService.GetCodeReviewFeedback(code)
		if err != nil {
			log.Printf("Error fetching feedback: %v", err)
			utils.RespondWithError(c, http.StatusInternalServerError, ErrProcessingReview)
			return
		}

		utils.RespondWithSuccess(c, gin.H{"feedback": feedback})
	}
}