package transcation

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func ResetAccountTable(db *gorm.DB) {
	if db.Debug().Migrator().HasTable(&Account{}) {
		db.Debug().Migrator().DropTable(&Account{})
	}

	db.AutoMigrate(&Account{})
}

func ResetTranscationTable(db *gorm.DB) {
	if db.Debug().Migrator().HasTable(&Transcation{}) {
		db.Debug().Migrator().DropTable(&Transcation{})
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
