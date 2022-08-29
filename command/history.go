package command

import (
	"atm-design/db"
	"atm-design/session"

	. "github.com/fatih/color"
)

func history(...string) (Message, error) {
	a := &db.Account{}
	if err := a.Find(session.GetSession().AccountID()); err != nil {
		return Message{RedString("Transaction failed. Try again later")}, err
	}

	tt, err := a.TransactionHistory()
	if err != nil {
		return Message{RedString("Transaction failed. Try again later")},
			err
	}

	if tt == nil || len(tt) == 0 {
		return Message{YellowString("No history found")},
			nil
	}

	var msg Message
	for _, t := range tt {
		msg = append(msg, GreenString("%s\n", t.String()))
	}
	return msg, nil
}
