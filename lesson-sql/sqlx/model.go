package sqlx

import "time"

type Employee struct {
	Id         uint
	Name       string
	Department string
	Salary     float32
}

type Book struct {
	Id          uint
	Catalog     string
	Title       string
	Price       float32
	Author      string
	Publisher   string
	PublishTime time.Time `db:"publish_time"`
}
