package transcation

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func ResetAccountTable(db *gorm.DB) {
	var record_count int
	db.Debug().Raw("SELECT COUNT(*) AS record_count FROM pg_class WHERE relname = 'accounts';").Scan(&record_count)
	if record_count == 1 {
		db.Debug().Exec("DROP TABLE accounts;")
	}

	db.AutoMigrate(&Account{})
}

func ResetTranscationTable(db *gorm.DB) {
	var record_count int
	db.Debug().Raw("SELECT COUNT(*) AS record_count FROM pg_class WHERE relname = 'transcations';").Scan(&record_count)
	if record_count == 1 {
		db.Debug().Exec("DROP TABLE transcations;")
	}

	db.AutoMigrate(&Transcation{})
}

func CreateAccount(db *gorm.DB, name string, balance uint) {
	db.Debug().Where("name = ?", name).Unscoped().Delete(&Account{})
	db.Create(&Account{Name: name, Balance: balance})
}

func Transfer(db *gorm.DB, transfer_from string, transfer_to string, amount uint) {
	tx := db.Begin()

	var from_account Account
	tx.Debug().Where("name = ?", transfer_from).First(&from_account)
	fmt.Println(from_account)

	var to_account Account
	tx.Debug().Where("name = ?", transfer_to).First(&to_account)
	fmt.Println(to_account)

	from_account_balance := from_account.Balance
	to_account_balance := to_account.Balance

	tx.SavePoint("sp1")
	tx.Debug().Create(&Transcation{FromAccountID: from_account.ID, ToAccountID: to_account.ID, Amount: amount})

	tx.Debug().Model(&from_account).Update("balance", strconv.Itoa(int(from_account_balance-amount)))
	tx.Debug().Model(&to_account).Update("balance", to_account_balance+amount)

	if from_account_balance < amount {
		tx.RollbackTo("sp1")
	} else {
		tx.Commit()
	}
}
