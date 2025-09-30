package gorm

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string
	PostCount uint
	Posts     []Post
}

type Post struct {
	gorm.Model
	Title    string
	Content  string
	Status   string
	Comments []Comment
	UserID   uint
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
}
