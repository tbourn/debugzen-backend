package utils

import "github.com/gin-gonic/gin"

func RespondWithSuccess(c *gin.Context, data gin.H) {
	c.JSON(200, data)
}

func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}