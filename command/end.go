package command

import (
	"os"

	"atm-design/session"
)

func end(...string) (Message, error) {
	session.GetSession().Stop()
	os.Exit(0)
	return Message{""}, nil
}
