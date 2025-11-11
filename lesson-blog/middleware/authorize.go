package middleware

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/henry-woo/golang-study/lesson-blog/controller"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.Request.FormValue("api_token")

		if token == "" {
			controller.RespondWithError(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			controller.RespondWithError(c, 401, "Invalid API token")
			return
		}

		c.Next()
	}
}
