package authorize

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/henry-woo/golang-study/lesson-blog/internal/store"

	"github.com/gin-gonic/gin"
)

func BearerFromHeader(c *gin.Context) string {
	h := c.GetHeader("Authorization")
	if strings.HasPrefix(h, "Bearer ") {
		return strings.TrimPrefix(h, "Bearer ")
	}

	return ""
}

func AuthorizeMiddleware(r *store.Redis) gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token, _ := c.Cookie("access_token")
		if access_token == "" {
			access_token = BearerFromHeader(c)
		}

		if access_token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		claims, err := ParseAccess(access_token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		context := context.Background()
		if _, err := r.GetUserByJTI(context, "access:"+claims.ID); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token revoked"})
			return
		}

		c.Set("userID", claims.Subject)

		c.Next()
	}
}

func MustCookie(c *gin.Context, name string) (string, error) {
	val, err := c.Cookie(name)
	if err != nil || val == "" {
		return "", errors.New("missing cookie: " + name)
	}
	return val, nil
}
