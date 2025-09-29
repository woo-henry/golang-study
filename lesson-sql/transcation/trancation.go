package transcation

import (
	"fmt"

	"gorm.io/gorm"
)

func CreateAccount(db *gorm.DB, name string, balance uint) {
	var record_count int
	db.Debug().Raw("SELECT COUNT(*) AS record_count FROM pg_class WHERE relname = 'accounts';").Scan(&record_count)
	if record_count == 0 {
		db.AutoMigrate(&Account{})
	}

	db.Debug().Where("name = ?", name).Delete(&Account{})
	db.Create(&Account{Name: name, Balance: balance})
}

func Transfer(db *gorm.DB, transfer_from string, transfer_to string, amount uint) {
	var record_count int
	db.Debug().Raw("SELECT COUNT(*) AS record_count FROM pg_class WHERE relname = 'transcations';").Scan(&record_count)
	if record_count == 0 {
		db.AutoMigrate(&Transcation{})
	}

	var from_account Account
	db.Debug().Where("name = ?", transfer_from).First(&from_account)
	fmt.Println(from_account)

	var to_account Account
	db.Debug().Where("name = ?", transfer_to).First(&to_account)

	fmt.Println(to_account)

	from_account.Balance -= amount
	to_account.Balance += amount
	db.Debug().Create(&Transcation{FromAccountID: from_account.ID, ToAccountID: to_account.ID, Amount: amount})
}
