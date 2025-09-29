package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGormDatabase() *gorm.DB {
	//dsn := "root:pass@tcp(127.0.0.1:3306)/golang-lesson?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=localhost user=golang-lesson password=golang-lesson dbname=golang-lesson port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
		panic("Failed to connect to database")
	}

	sql_db, err := db.DB()
	if err != nil {
		log.Fatal("Failed to open db")
		panic("Failed to open db")
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sql_db.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sql_db.SetMaxOpenConns(20)

	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sql_db.SetConnMaxLifetime(time.Hour)

	return db
}

func AutoMigrate(db *gorm.DB, dst ...interface{}) error {
	return db.Migrator().AutoMigrate(dst...)
}
