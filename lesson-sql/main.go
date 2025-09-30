package main

import (
	"fmt"

	"github.com/henry-woo/golang-study/lesson-sql/crud"
	"github.com/henry-woo/golang-study/lesson-sql/database"
	"github.com/henry-woo/golang-study/lesson-sql/gorm"
	"github.com/henry-woo/golang-study/lesson-sql/sqlx"
	"github.com/henry-woo/golang-study/lesson-sql/transcation"
)

func main() {

	// lesson 1 ~ 2
	db_gorm := database.InitGormDatabase()
	{

		crud.ResetStudentTable(db_gorm)

		// lesson 1
		crud.CreateStudent(db_gorm, "赵大", 10, "一年级")
		crud.CreateStudent(db_gorm, "钱二", 12, "二年级")
		crud.CreateStudent(db_gorm, "张三", 20, "三年级")
		crud.CreateStudent(db_gorm, "李四", 22, "四年级")
		crud.CreateStudent(db_gorm, "王五", 23, "五年级")
		crud.CreateStudent(db_gorm, "谢六", 24, "六年级")
		students := crud.QueryStudents(db_gorm, 18)
		student_size := len(students)
		if student_size > 0 {
			fmt.Println(students)
		} else {
			fmt.Println("Not Found Over 18 Years Old Students")
		}

		updated := crud.UpdateStudentGrade(db_gorm, "张三", "四年级")
		fmt.Println("update student row : ", updated)

		crud.DeleteStudents(db_gorm, 15)

		// lesson 2
		transcation.ResetAccountTable(db_gorm)
		transcation.CreateAccount(db_gorm, "A", 100)
		transcation.CreateAccount(db_gorm, "B", 100)

		transcation.ResetTranscationTable(db_gorm)
		transcation.Transfer(db_gorm, "A", "B", 100)
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

	// lesson 5 ~ 7
	{
		// lesson 5
		gorm.ResetUserTable(db_gorm)
		gorm.ResetPostTable(db_gorm)
		gorm.ResetCommentTable(db_gorm)

		// lesson 6.1
		posts := gorm.QueryUserPosts(db_gorm, 1)
		post_size := len(posts)
		if post_size > 0 {
			fmt.Println("Found Post Size = ", post_size)
			fmt.Println("Found Post = ", posts[0])
		} else {
			fmt.Println("Not Found User Posts")
		}

		// lesson 6.2
		post := gorm.QueryMaxCommentsPost(db_gorm)
		fmt.Println(post)

		// lesson 7.1
		gorm.CreateUserPostWithHook(db_gorm, "测试", "测试", 3)

		// lesson 7.2
		gorm.DeleteCommentWithHook(db_gorm, 1)
	}
}
