package command

import (
	"fmt"

	"atm-design/db"
	"atm-design/session"

	. "github.com/fatih/color"
)

func authorize(args ...string) (Message, error) {
	s := session.GetSession()

	if s.IsAuthorized() {
		return Message{YellowString("Already authorized.")},
			fmt.Errorf("already authorized")
	}

	a := &db.Account{}
	err := a.Authorize(args[0], args[1])
	if err != nil {
		return Message{YellowString("Authorization failed.")},
			fmt.Errorf("authorization failed")
	}

	go s.Start(a.AccountID)

	return Message{GreenString("%s successfully authorized.", a)}, nil
}
