package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"atm-design/command"
	"atm-design/db"

	"github.com/fatih/color"
)

func TestAll(t *testing.T) {
	var tests = []struct {
		cmd  string
		want []string
	}{
		{"bad command", []string{color.RedString("Unknown command 'bad'.")}},
		{"authorize 2859459814 0000", []string{color.YellowString("Authorization failed.")}},
		{"authorize 2859459814 7386", []string{color.GreenString("2859459814 successfully authorized.")}},
		{"authorize 2859459814 7386", []string{color.YellowString("Already authorized.")}},
		{"withdraw -1", []string{
			color.RedString("Incorrect amount, must be greater than 0."),
		}},
		{"withdraw 10", []string{color.YellowString("Withdrawal amount must be a multiple of 20.")}},
		{"withdraw 100", []string{
			color.GreenString("Amount dispensed: $100.00"),
			color.YellowString("You have been charged an overdraft fee of $5.00."),
			color.GreenString("Current balance: -94.76"),
		}},
		{"withdraw 20", []string{color.YellowString("Your account is overdrawn! You may not make withdrawals at this time.")}},
		{"deposit 100", []string{color.GreenString("Current balance: 5.24")}},
		{"withdraw 1000", []string{
			color.GreenString("Amount dispensed: $1,000.00"),
			color.YellowString("You have been charged an overdraft fee of $5.00."),
			color.GreenString("Current balance: -999.76"),
		}},
		{"deposit 10000", []string{color.GreenString("Current balance: 9,000.24")}},
		{"withdraw 10000", []string{
			color.YellowString("Unable to dispense full amount requested at this time."),
			color.GreenString("Amount dispensed: $8,900.00"),
			color.GreenString("Current balance: 100.24"),
		}},
		{"withdraw 200", []string{
			color.YellowString("Unable to process your withdrawal at this time."),
		}},
		{"logout", []string{color.GreenString("Account 2859459814 logged out.")}},
		{"logout", []string{color.YellowString("No account is currently authorized.")}},
	}

	dbFile := "test/" + t.Name() + ".sqlite"
	defer os.Remove(dbFile)
	_, err := db.Connect(dbFile)
	if err != nil {
		t.Fatalf(`%v`, err)
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.cmd)
		t.Run(testName, func(t *testing.T) {
			time.Sleep(100 * time.Millisecond)
			msg, err := command.Executor(strings.Split(tt.cmd, " ")...)
			for i := range msg {
				if msg[i] != tt.want[i] {
					t.Errorf("got %s, want %s", msg[i], tt.want[i])
					t.Logf("error: %v", err)
				}
			}
		})
	}
}
