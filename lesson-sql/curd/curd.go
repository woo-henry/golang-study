package curd

import (
	"gorm.io/gorm"
)

func CreateStudent(db *gorm.DB, name string, age uint, grade string) {
	var record_count int
	db.Debug().Raw("SELECT COUNT(*) AS record_count FROM pg_class WHERE relname = 'students';").Scan(&record_count)
	if record_count == 0 {
		db.Debug().Exec("DROP TABLE students;")
	}

	db.Debug().AutoMigrate(&Student{})

	db.Debug().Create(&Student{Name: name, Age: age, Grade: grade})
}
