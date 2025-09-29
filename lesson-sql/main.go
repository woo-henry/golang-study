package main

import (
	"fmt"

	"github.com/henry-woo/golang-study/lesson-sql/curd"
	"github.com/henry-woo/golang-study/lesson-sql/database"
	"github.com/henry-woo/golang-study/lesson-sql/transcation"
)

func main() {

	// lesson 1 ~ 2
	{
		db_gorm := database.InitGormDatabase()
		curd.ResetStudentTable(db_gorm)

		// lesson 1
		curd.CreateStudent(db_gorm, "赵大", 10, "一年级")
		curd.CreateStudent(db_gorm, "钱二", 12, "二年级")
		curd.CreateStudent(db_gorm, "张三", 20, "三年级")
		curd.CreateStudent(db_gorm, "李四", 22, "四年级")
		curd.CreateStudent(db_gorm, "王五", 23, "五年级")
		curd.CreateStudent(db_gorm, "谢六", 24, "六年级")
		students := curd.QueryStudents(db_gorm, 18)
		student_size := len(students)
		if student_size > 0 {
			fmt.Println(students)
		} else {
			fmt.Println("Not Found Over 18 Years Old Students")
		}

		updated := curd.UpdateStudentGrade(db_gorm, "张三", "四年级")
		fmt.Println("update student row : {}", updated)

		curd.DeleteStudents(db_gorm, 15)

		// lesson 2
		transcation.ResetAccountTable(db_gorm)
		transcation.CreateAccount(db_gorm, "A", 100)
		transcation.CreateAccount(db_gorm, "B", 100)

		transcation.ResetTranscationTable(db_gorm)
		transcation.Transfer(db_gorm, "A", "B", 100)
		transcation.Transfer(db_gorm, "A", "B", 100)
	}

	// lesson 3 ~ 4
	/*
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
	*/
}
