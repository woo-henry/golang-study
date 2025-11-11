package controller

import (
	"github.com/gin-gonic/gin"
)

func RespondWithSuccess(c *gin.Context, code int, message interface{}) {
	c.JSON(code, gin.H{"message": message})
}

func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
