package service

import (
	"github.com/henry-woo/golang-study/lesson-blog/model"
)

func CreateUser(user *model.User) (*model.User, bool) {
	return user, true
}

func ExistUser(user *model.User) (*model.User, bool) {
	return nil, true
}
