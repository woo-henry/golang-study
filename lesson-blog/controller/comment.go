package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henry-woo/golang-study/lesson-blog/model"
	"github.com/henry-woo/golang-study/lesson-blog/service"
)

func CommentCreate(c *gin.Context) {
	var in struct {
		UserID  uint   `json:"user_id"`
		PostID  uint   `json:"post_id"`
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

	comment := model.Comment{}
	comment.Content = in.Content
	comment.PostID = in.PostID
	comment.UserID = in.UserID
	if comment.UserID == 0 {
		comment.UserID = userID
	}

	created_comment, created := service.CreateComment(&comment)
	if created_comment == nil || !created {
		RespondWithError(c, http.StatusInternalServerError, "failed to create comment")
		return
	}

	RespondWithSuccess(c, http.StatusOK, created_comment)
}

func CommentQuery(c *gin.Context) {
	post_id := c.GetUint("post_id")
	if post_id == 0 {
		RespondWithError(c, http.StatusBadRequest, "invalid parameter")
		return
	}

	comments, success := service.QueryComments(post_id)
	if !success {
		RespondWithError(c, http.StatusInternalServerError, "failed paramto query comments")
		return
	}

	RespondWithSuccess(c, http.StatusOK, comments)
}
