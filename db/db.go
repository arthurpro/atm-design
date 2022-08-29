package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Connect(name string) (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(sqlite.Open(name), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	err = migrate()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Instance() *gorm.DB {
	return db
}

func migrate() error {
	err := db.AutoMigrate(
		&Setting{},
		&Account{},
		&Transaction{},
	)
	if err != nil {
		return err
	}

	if err := db.
		Where("account_id = ?", "2859459814").
		First(&Account{}).Error; err != nil {
		var accounts = []Account{
			{2859459814, "7386", 10.24},
			{1434597300, "4557", 90000.55},
			{7089382418, "0075", 0.00},
			{2001377812, "5950", 60.00},
		}
		if err := db.Create(&accounts).Error; err != nil {
			return err
		}

		var settings = []Setting{
			{"balance", 10000},
			{"overdraft_fee", 5},
			{"bill_value", 20},
		}
		if err := db.Create(&settings).Error; err != nil {
			return err
		}
	}
	return nil
}
