package command

import (
	"fmt"
	"strconv"

	"atm-design/db"
	"atm-design/session"

	. "github.com/fatih/color"
)

func deposit(args ...string) (Message, error) {
	amount, err := strconv.ParseFloat(args[0], 10)
	if err != nil || amount < 0.01 {
		return Message{RedString("Incorrect amount, must be greater than 0.00")},
			fmt.Errorf("incorrect deposit amount '%f'", amount)
	}

	a := &db.Account{}
	if err := a.Find(session.GetSession().AccountID()); err != nil {
		return Message{RedString("Transaction failed. Try again later")}, err
	}

	if _, err = a.RecordTransaction(amount); err != nil {
		return Message{RedString("Transaction failed. Try again later")}, err
	}

	return Message{GreenString("Current balance: %s", a.HumanBalance())}, nil
}
