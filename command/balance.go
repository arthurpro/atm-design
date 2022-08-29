package command

import (
	"atm-design/db"
	"atm-design/session"

	. "github.com/fatih/color"
)

func balance(...string) (Message, error) {
	a := &db.Account{}
	if err := a.Find(session.GetSession().AccountID()); err != nil {
		return Message{RedString("Transaction failed. Try again later")}, err
	}
	return Message{GreenString("Current balance: %s", a.HumanBalance())}, nil
}
