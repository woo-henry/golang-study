package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/henry-woo/golang-study/lesson-blog/internal/authorize"
	"github.com/henry-woo/golang-study/lesson-blog/internal/store"
	"github.com/henry-woo/golang-study/lesson-blog/model"
	"github.com/henry-woo/golang-study/lesson-blog/service"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var in struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&in); err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, "failed to hash password")
		return
	}

	user := model.User{}
	user.Username = in.Username
	user.Password = string(hashedPassword)

	createdUser, created := service.CreateUser(&user)
	if createdUser == nil || !created {
		RespondWithError(c, http.StatusInternalServerError, "failed to create user")
		return
	}

	RespondWithSuccess(c, http.StatusOK, "user register successfully")
}

func Login(c *gin.Context) {
	var in struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&in); err != nil {
		RespondWithError(c, http.StatusUnauthorized, "invalid request")
		return
	}

	existUser, exist := service.FindUserByUsername(in.Username)
	if existUser == nil || !exist {
		RespondWithError(c, http.StatusUnauthorized, "user not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(in.Password)); err != nil {
		RespondWithError(c, http.StatusUnauthorized, "invalid username or password")
		return
	}

	toks, err := authorize.IssueTokens(strconv.Itoa(int(existUser.ID)))
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, "could not issue tokens")
		return
	}

	redis_client := store.RedisClient()
	if err := authorize.Persist(c, redis_client, toks); err != nil {
		RespondWithError(c, http.StatusInternalServerError, "could not persist tokens")
		return
	}

	authorize.SetAuthCookies(c, toks)

	RespondWithSuccess(c, http.StatusOK, toks)
}

func Logout(c *gin.Context) {
	access_token, _ := c.Cookie("access_token")
	refresh_token, _ := c.Cookie("refresh_token")
	context := context.Background()
	redis_client := store.RedisClient()

	if access_token != "" {
		if claims, err := authorize.ParseAccess(access_token); err == nil {
			_ = redis_client.DelJTI(context, "access:"+claims.ID)
		}
	}

	if refresh_token != "" {
		if claims, err := authorize.ParseRefresh(refresh_token); err == nil {
			_ = redis_client.DelJTI(context, "refresh:"+claims.ID)
		}
	}

	authorize.ClearAuthCookies(c)

	RespondWithSuccess(c, http.StatusOK, "user logout success")
}
