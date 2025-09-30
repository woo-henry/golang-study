package crud

import (
	"gorm.io/gorm"
)

func ResetStudentTable(db *gorm.DB) {
	if db.Debug().Migrator().HasTable(&Student{}) {
		db.Debug().Migrator().DropTable(&Student{})
	}

	db.Debug().AutoMigrate(&Student{})
}

func CreateStudent(db *gorm.DB, name string, age uint, grade string) {
	db.Debug().Create(&Student{Name: name, Age: age, Grade: grade})
}

func QueryStudents(db *gorm.DB, age uint) []Student {
	var students []Student
	db.Debug().Where("age > ?", age).Find(&students)
	return students
}

func UpdateStudentGrade(db *gorm.DB, name string, grade string) int64 {
	result := db.Debug().Model(&Student{}).Where("name = ?", name).Update("grade", grade)
	return result.RowsAffected
}

func DeleteStudents(db *gorm.DB, age uint) {
	db.Debug().Where("age < ?", age).Unscoped().Delete(&Student{})
}
