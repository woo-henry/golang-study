package service

import (
	"github.com/henry-woo/golang-study/lesson-blog/database"
	"github.com/henry-woo/golang-study/lesson-blog/model"
)

func CreateComment(comment *model.Comment) (*model.Comment, bool) {
	db := database.GetInstance()

	var err = db.Debug().Model(&model.Comment{}).Create(comment).Take(comment).Error
	if err != nil {
		return nil, false
	}

	return comment, true
}

func QueryComments(post_id uint) ([]model.Comment, bool) {
	var result []model.Comment

	db := database.GetInstance()

	var err = db.Debug().Model(&model.Comment{}).Where("post_id = ?", post_id).Find(&result).Error
	if err != nil {
		return result, false
	}

	return result, true
}
