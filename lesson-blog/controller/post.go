package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henry-woo/golang-study/lesson-blog/model"
	"github.com/henry-woo/golang-study/lesson-blog/service"
)

func PostCreate(c *gin.Context) {
	var in struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&in); err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := GetUserID(c)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "no authorized")
		return
	}

	post := model.Post{}
	post.Title = in.Title
	post.Content = in.Content
	post.UserID = userID

	created_post, created := service.CreatePost(&post)
	if created_post == nil || !created {
		RespondWithError(c, http.StatusInternalServerError, "failed to create post")
		return
	}

	RespondWithSuccess(c, http.StatusOK, created_post)
}

func PostRemove(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := GetUserID(c)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "no authorized")
		return
	}

	if post.UserID > 0 && post.UserID != userID {
		RespondWithError(c, http.StatusBadRequest, "not allow remove")
		return
	} else {
		post.UserID = userID
	}

	service.RemovePost(post.ID)

	RespondWithSuccess(c, http.StatusOK, "post remove success")
}

func PostUpdate(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := GetUserID(c)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "no authorized")
		return
	}

	if post.UserID > 0 && post.UserID != userID {
		RespondWithError(c, http.StatusBadRequest, "not allow update")
		return
	} else {
		post.UserID = userID
	}

	service.UpdatePost(&post)

	RespondWithSuccess(c, http.StatusOK, post)
}

func PostQuery(c *gin.Context) {
	post_id := c.GetUint("post_id")
	if post_id > 0 {
		post := service.QueryPost(post_id)
		RespondWithSuccess(c, http.StatusOK, post)
	} else {
		posts := service.QueryPosts()
		RespondWithSuccess(c, http.StatusOK, posts)
	}
}
