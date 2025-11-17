package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/henry-woo/golang-study/lesson-blog/controller"
	"github.com/henry-woo/golang-study/lesson-blog/database"
	"github.com/henry-woo/golang-study/lesson-blog/internal/authorize"
	"github.com/henry-woo/golang-study/lesson-blog/internal/store"
	"github.com/henry-woo/golang-study/lesson-blog/model"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	for _, k := range []string{"ACCESS_SECRET", "REFRESH_SECRET"} {
		if os.Getenv(k) == "" {
			log.Fatalf("%s not set", k)
		}
	}

	db := database.GetInstance()
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})

	rds := store.RedisClient()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	public := router.Group("/public")
	public.POST("/register", controller.Register)
	public.POST("/login", controller.Login)
	public.POST("/logout", authorize.AuthorizeMiddleware(rds), controller.Logout)
	public.POST("/token/refresh", controller.RefreshToken)

	protected := router.Group("/api")
	protected.Use(authorize.AuthorizeMiddleware(rds))
	protected.POST("/post/create", authorize.AuthorizeMiddleware(rds), controller.PostCreate)
	protected.POST("/post/delete", authorize.AuthorizeMiddleware(rds), controller.PostRemove)
	protected.POST("/post/update", authorize.AuthorizeMiddleware(rds), controller.PostUpdate)
	protected.GET("/post/query", authorize.AuthorizeMiddleware(rds), controller.PostQuery)
	protected.POST("/comment/create", authorize.AuthorizeMiddleware(rds), controller.CommentCreate)
	protected.GET("/comment/query", authorize.AuthorizeMiddleware(rds), controller.CommentQuery)

	log.Fatal(router.Run(":8080"))
}
