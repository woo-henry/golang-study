package sqlx

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func QueryDepartmentEmployees(db *sqlx.DB, department string) []Employee {
	sql := "select id, name, department, salary from employees where department = $1"
	var employees []Employee
	err := db.Select(&employees, sql, department)
	if err != nil {
		log.Println(err.Error())
	}

	return employees
}

func QueryMaxSalaryEmployee(db *sqlx.DB) Employee {
	sql := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1;"
	var employee Employee
	err := db.Get(&employee, sql)
	if err != nil {
		log.Println(err.Error())
	}
	return employee
}

func CreateEmployees(db *sqlx.DB) {
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM employees")
	tx.MustExec("INSERT INTO employees (id, name, department, salary) VALUES ($1, $2, $3, $4)", 1, "Jason", "技术部", 15000.23)
	tx.MustExec("INSERT INTO employees (id, name, department, salary) VALUES ($1, $2, $3, $4)", 2, "John", "市场部", 12000.34)
	tx.MustExec("INSERT INTO employees (id, name, department, salary) VALUES ($1, $2, $3, $4)", 3, "Herny", "技术部", 18000.45)
	tx.NamedExec("INSERT INTO employees (id, name, department, salary) VALUES (:id, :name, :department, :salary)", &Employee{4, "Susan", "行政部", 10000.56})
	tx.NamedExec("INSERT INTO employees (id, name, department, salary) VALUES (:id, :name, :department, :salary)", &Employee{5, "Bill", "技术部", 8000.99})
	tx.Commit()
}

func QueryBestsellingBooks(db *sqlx.DB, price float32) []Book {
	sql := "select id, catalog, name, price, author, publisher, publish_time from books where price > $1"
	var books []Book
	err := db.Select(&books, sql, price)
	if err != nil {
		log.Println(err.Error())
	}

	return books
}

func CreateBooks(db *sqlx.DB) {
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM books")
	tx.MustExec("INSERT INTO books (id, catalog, name, price, author, publisher, publish_time) VALUES ($1, $2, $3, $4, $5, $6, $7)", 1, "哲学", "亚里士多德是个话痨", 66.23, "亚里士多德", "古希腊出版社", time.Now())
	tx.MustExec("INSERT INTO books (id, catalog, name, price, author, publisher, publish_time) VALUES ($1, $2, $3, $4, $5, $6, $7)", 2, "商业", "风口来了吗？", 55.34, "雷布斯", "小米出版社", time.Now())
	tx.MustExec("INSERT INTO books (id, catalog, name, price, author, publisher, publish_time) VALUES ($1, $2, $3, $4, $5, $6, $7)", 3, "科技", "Golang 实践与应用", 44.45, "说英雄是英雄", "微软出版社", time.Now())
	tx.NamedExec("INSERT INTO books (id, catalog, name, price, author, publisher, publish_time) VALUES (:id, :catalog, :name, :price, :author, :publisher, :publish_time)", &Book{4, "艺术", "眉毛上的蚂蚁", 33.56, "哈拉", "WhatApp出版社", time.Now()})
	tx.NamedExec("INSERT INTO books (id, catalog, name, price, author, publisher, publish_time) VALUES (:id, :catalog, :name, :price, :author, :publisher, :publish_time)", &Book{5, "教科书", "葡萄牙语入门教材", 22.99, "飞哥", "北京胡同出版社", time.Now()})
	tx.Commit()
}
