package transcation

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name    string
	Balance uint
}

type Transcation struct {
	FromAccountID uint `gorm:"references:id"`
	ToAccountID   uint `gorm:"references:id"`
	Amount        uint
}
