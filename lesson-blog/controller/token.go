package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henry-woo/golang-study/lesson-blog/internal/authorize"
	"github.com/henry-woo/golang-study/lesson-blog/internal/store"
)

func RefreshToken(c *gin.Context) {
	refresh_token, err := authorize.MustCookie(c, "refresh_token")
	if err != nil {
		RespondWithError(c, http.StatusUnauthorized, "missing refresh token")
		return
	}

	claims, err := authorize.ParseRefresh(refresh_token)
	if err != nil {
		RespondWithError(c, http.StatusUnauthorized, "invalid refresh token")
		return
	}

	redis_client := store.RedisClient()
	ctx := context.Background()
	if _, err := redis_client.GetUserByJTI(ctx, "refresh:"+claims.ID); err != nil {
		RespondWithError(c, http.StatusUnauthorized, "refresh revoked")
		return
	}

	_ = redis_client.DelJTI(ctx, "refresh:"+claims.ID)

	toks, err := authorize.IssueTokens(claims.Subject)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, gin.H{"error": "could not issue new tokens"})
		return
	}

	if err := authorize.Persist(ctx, redis_client, toks); err != nil {
		RespondWithError(c, http.StatusInternalServerError, gin.H{"error": "could not persist new tokens"})
		return
	}

	authorize.SetAuthCookies(c, toks)

	RespondWithSuccess(c, http.StatusOK, "success")
}
