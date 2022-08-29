package db

import (
	"fmt"
	"math"
	"time"

	"github.com/dustin/go-humanize"
)

type Transaction struct {
	TransactionID uint `gorm:"autoIncrement"`
	AccountID     uint
	Account       Account `gorm:"references:AccountID"`
	DateTime      *time.Time
	Amount        float64
	Balance       float64
}

func (t *Transaction) HumanAmount() string {
	return humanize.FormatFloat("", t.Amount)
}

func (t *Transaction) HumanAmountAbs() string {
	return humanize.FormatFloat("", math.Abs(t.Amount))
}

func (t *Transaction) HumanBalance() string {
	return humanize.FormatFloat("", t.Balance)
}

func (t *Transaction) HumanDateTime() string {
	return t.DateTime.Format("2006-02-01 15:04:05")
}

func (t *Transaction) String() string {
	return fmt.Sprintf(
		"%s %s %s",
		t.HumanDateTime(),
		t.HumanAmount(),
		t.HumanBalance(),
	)
}
