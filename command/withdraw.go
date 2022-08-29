package command

import (
	"fmt"
	"strconv"

	"atm-design/db"
	"atm-design/session"

	. "github.com/fatih/color"
)

func withdraw(args ...string) (Message, error) {
	amountInt, err := strconv.Atoi(args[0])
	amount := float64(amountInt)
	if err != nil || amountInt <= 0 {
		return Message{RedString("Incorrect amount, must be greater than 0.")},
			fmt.Errorf("incorrect amount for withdrawal '%s'", args[0])
	}

	a := &db.Account{}
	if err := a.Find(session.GetSession().AccountID()); err != nil {
		return Message{RedString("Transaction failed. Try again later")}, err
	}

	if a.Balance <= 0 {
		return Message{YellowString("Your account is overdrawn! You may not make withdrawals at this time.")},
			fmt.Errorf("account is overdrawn")
	}

	s := &db.Setting{}
	var bill int
	if v, err := s.BillValue(); err == nil {
		bill = int(v)
	}

	if bill < 1 {
		return Message{YellowString("Unable to process your withdrawal at this time.")},
			fmt.Errorf("bill value is not available")
	}

	if (amountInt % bill) != 0 {
		return Message{YellowString("Withdrawal amount must be a multiple of %s.", strconv.Itoa(bill))},
			fmt.Errorf("wrong amount for withdrawal")
	}

	s = &db.Setting{}
	var balance float64
	var msg Message
	if balance, err = s.Balance(); err == nil {
		if int64(balance) < int64(bill) {
			return Message{YellowString("Unable to process your withdrawal at this time.")},
				fmt.Errorf("atm is out of bills")
		}
		if balance < amount {
			amount = balance
			msg = append(msg, YellowString("Unable to dispense full amount requested at this time."))
		}
	}
	if err = s.UpdateBalance(balance - amount); err != nil {
		msg = append(msg, YellowString("Unable to process your withdrawal at this time."))
		return msg, err
	}

	h, err := a.RecordTransaction(-amount)
	if err != nil {
		return Message{RedString("Transaction failed. Try again later")}, err
	}

	msg = append(msg, GreenString("Amount dispensed: $%s", h.HumanAmountAbs()))

	if a.Balance < 0 {
		s := &db.Setting{}
		if fee, err := s.OverdraftFee(); err == nil {
			h, _ := a.RecordTransaction(-fee)
			msg = append(msg, YellowString("You have been charged an overdraft fee of $%s.", h.HumanAmountAbs()))
		}
	}
	msg = append(msg, GreenString("Current balance: %s", a.HumanBalance()))
	return msg, nil
}
