package main

import (
	"github.com/henry-woo/golang-study/lesson-sql/curd"
	"github.com/henry-woo/golang-study/lesson-sql/database"
	"github.com/henry-woo/golang-study/lesson-sql/transcation"
)

func main() {
	db := database.InitDatabase()

	curd.CreateStudent(db, "张三", 20, "三年级")

	transcation.CreateAccount(db, "A", 1000)
	transcation.CreateAccount(db, "B", 1000)
	transcation.Transfer(db, "A", "B", 100)
}
