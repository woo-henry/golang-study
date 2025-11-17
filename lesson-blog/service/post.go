package service

import (
	"github.com/henry-woo/golang-study/lesson-blog/database"
	"github.com/henry-woo/golang-study/lesson-blog/model"
)

func CreatePost(post *model.Post) (*model.Post, bool) {
	db := database.GetInstance()

	var err = db.Debug().Model(&model.Post{}).Create(post).Error
	if err != nil {
		return nil, false
	}

	return post, true
}

func RemovePost(post_id uint) bool {
	db := database.GetInstance()

	var err = db.Debug().Where("id = ?", post_id).Delete(&model.Post{})
	if err != nil {
		return false
	}

	return true
}

func UpdatePost(post *model.Post) (*model.Post, bool) {
	db := database.GetInstance()

	err := db.Debug().Model(&model.Post{}).Where("id = ?", post.ID).Save(post).Take(post).Error
	if err != nil {
		return post, false
	}

	return post, true
}

func QueryPosts() []model.Post {
	var result []model.Post

	db := database.GetInstance()
	var err = db.Debug().Model(&model.Post{}).Find(&result).Error
	if err != nil {
		return result
	}

	return result
}

func QueryPost(post_id uint) *model.Post {
	var result model.Post

	db := database.GetInstance()
	var err = db.Debug().Where("id = ?", post_id).Take(&result).Error
	if err != nil {
		return nil
	}

	return &result
}
