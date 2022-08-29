package db

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
)

type Account struct {
	AccountID uint `gorm:"primaryKey"`
	Pin       string
	Balance   float64
}

func (a *Account) Find(acc any) error {
	tx := db.Where("account_id = ?", acc).First(&a)
	return tx.Error
}

func (a *Account) Authorize(acc, pin string) error {
	tx := db.Where("account_id = ? and pin = ?", acc, pin).First(&a)
	if tx.Error != nil {
		return fmt.Errorf("authorization failed")
	}
	return nil
}

func (a *Account) HumanBalance() string {
	return humanize.FormatFloat("", a.Balance)
}

func (a *Account) UpdateBalance(amount float64) error {
	a.Balance = amount
	tx := db.Save(a)
	return tx.Error
}

func (a *Account) String() string {
	return fmt.Sprintf("%d", a.AccountID)
}

func (a *Account) RecordTransaction(amount float64) (*Transaction, error) {
	now := time.Now()
	t := &Transaction{
		TransactionID: 0,
		AccountID:     a.AccountID,
		Account:       *a,
		DateTime:      &now,
		Amount:        amount,
		Balance:       a.Balance + amount,
	}
	tx := db.Create(t)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := a.UpdateBalance(t.Balance); err != nil {
		return nil, err
	}

	return t, nil
}

func (a *Account) TransactionHistory() ([]*Transaction, error) {
	var tt []*Transaction

	tx := db.
		Where("account_id = ?", a.AccountID).
		Order("date_time desc").
		Find(&tt)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return tt, nil
}
