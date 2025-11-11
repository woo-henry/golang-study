package controller

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/henry-woo/golang-study/lesson-blog/model"
	"github.com/henry-woo/golang-study/lesson-blog/service"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}
	user.Password = string(hashedPassword)

	createdUser, created := service.CreateUser(&user)
	if createdUser == nil || !created {
		RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	RespondWithSuccess(c, http.StatusCreated, "User registered successfully")
}

func UserLogin(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	existUser, exist := service.ExistUser(&user)
	if existUser == nil || !exist {
		RespondWithError(c, http.StatusUnauthorized, "User not found")
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password)); err != nil {
		RespondWithError(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       existUser.ID,
		"username": existUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	RespondWithSuccess(c, http.StatusCreated, tokenString)
}
