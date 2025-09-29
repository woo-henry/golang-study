package crud

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name  string
	Age   uint
	Grade string
}
