package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB = nil
	mutex    sync.Mutex
)

func GetInstance() *gorm.DB {
	mutex.Lock()
	if instance == nil {
		instance = OpenDatabase()
	}
	mutex.Unlock()

	return instance
}

func OpenDatabase() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", DbHost, DbPort, DbName, DbUser, DbPassword)
	//dsn := "host=localhost user=golang-lesson password=golang-lesson dbname=golang-lesson port=5432 sslmode=disable TimeZone=Asia/Shanghai"
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
