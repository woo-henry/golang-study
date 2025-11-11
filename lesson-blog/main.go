package main

import (
	"github.com/gin-gonic/gin"
	"github.com/henry-woo/golang-study/lesson-blog/controller"
	"github.com/henry-woo/golang-study/lesson-blog/database"
	"github.com/henry-woo/golang-study/lesson-blog/middleware"
	"github.com/henry-woo/golang-study/lesson-blog/model"
)

func main() {

	db_gorm := database.InitGormDatabase()
	db_gorm.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})

	router := gin.Default()
	router.Use(middleware.TokenAuthMiddleware())

	router.POST("/register", controller.UserRegister)
	router.POST("/login", controller.UserLogin)

	router.POST("/post/create", controller.PostCreate)
	router.POST("/post/delete", controller.PostRemove)
	router.POST("/post/update", controller.PostUpdate)
	router.GET("/post/query", controller.PostQuery)

	router.POST("/comment/create", controller.CommentCreate)
	router.GET("/comment/query", controller.CommentQuery)

	router.Run() // default listen on 0.0.0.0:8080
}
