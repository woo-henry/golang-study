package main

import (
	"fmt"

	"github.com/henry-woo/golang-study/lesson-sql/curd"
	"github.com/henry-woo/golang-study/lesson-sql/database"
	"github.com/henry-woo/golang-study/lesson-sql/sqlx"
	"github.com/henry-woo/golang-study/lesson-sql/transcation"
)

func main() {

	// lesson 1 ~ 2
	{
		db_gorm := database.InitGormDatabase()

		// lesson 1
		curd.CreateStudent(db_gorm, "张三", 20, "三年级")

		// lesson 2
		transcation.CreateAccount(db_gorm, "A", 1000)
		transcation.CreateAccount(db_gorm, "B", 1000)
		transcation.Transfer(db_gorm, "A", "B", 100)
	}

	// lesson 3 ~ 4
	{
		db_sqlx := database.InitSqlxDatabase()
		sqlx.CreateEmployees(db_sqlx) // create employee test data

		// lesson 3.1
		employees := sqlx.QueryDepartmentEmployees(db_sqlx, "技术部")
		employees_size := len(employees)
		if employees_size > 0 {
			fmt.Println(employees)
		} else {
			fmt.Println("No Employees")
		}

		// lesson 3.2
		employee := sqlx.QueryMaxSalaryEmployee(db_sqlx)
		fmt.Println(employee)

		// lesson 4
		sqlx.CreateBooks(db_sqlx) // create book test data
		books := sqlx.QueryBestsellingBooks(db_sqlx, 50)
		book_size := len(books)
		if book_size > 0 {
			fmt.Println(books)
		} else {
			fmt.Println("No Books")
		}
	}
}
