package review

import (
	"log"
	"net/http"
	"strings"

	"github.com/tbourn/debugzen-backend/internal/utils"
	"github.com/tbourn/debugzen-backend/services"

	"github.com/gin-gonic/gin"
)

type ReviewRequest struct {
	Code string `json:"code" binding:"required"`
}

// ReviewHandler handles code review requests
type ReviewHandler struct {
	Service *services.ReviewService
}

// NewReviewHandler initializes a new ReviewHandler
func NewReviewHandler(reviewService *services.ReviewService) *ReviewHandler {
    return &ReviewHandler{Service: reviewService}
}

// @Summary Submit code for review
// @Description Sends code to OpenAI for review and feedback
// @Tags review
// @Accept  json
// @Produce  json
// @Param code body ReviewRequest true "Code to analyze"
// @Success 200 {object} services.ReviewResponse
// @Failure 400 {object} map[string]string
// @Router /review [post]
func (h *ReviewHandler) Review(c *gin.Context) {
	var req ReviewRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	code := strings.TrimSpace(req.Code)
	if code == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Code input is required")
		return
	}

	feedback, err := h.Service.GetCodeReviewFeedback(code)
	if err != nil {
		log.Printf("Error fetching feedback: %v", err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to process code review")
		return
	}

	c.JSON(http.StatusOK, services.ReviewResponse{Feedback: feedback})
}