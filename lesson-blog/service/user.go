package service

import (
	"github.com/henry-woo/golang-study/lesson-blog/database"
	"github.com/henry-woo/golang-study/lesson-blog/model"
)

func CreateUser(user *model.User) (*model.User, bool) {
	db := database.GetInstance()

	var err = db.Debug().Create(&user).Error
	if err != nil {
		return nil, false
	}

	return user, true
}

func FindUserByUsername(username string) (*model.User, bool) {
	db := database.GetInstance()

	u := model.User{}
	err := db.Debug().Model(model.User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return nil, false
	}

	return &u, true
}
