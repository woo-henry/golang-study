package main

import (
	"log"

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

	public := router.Group("/public")
	public.POST("/register", controller.UserRegister)
	public.POST("/login", controller.UserLogin)

	api := router.Group("/api/v1")
	api.POST("/post/create", controller.PostCreate)
	api.POST("/post/delete", controller.PostRemove)
	api.POST("/post/update", controller.PostUpdate)
	api.GET("/post/query", controller.PostQuery)
	api.POST("/comment/create", controller.CommentCreate)
	api.GET("/comment/query", controller.CommentQuery)

	log.Fatal(router.Run(":8080"))
}
