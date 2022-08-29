package command

import (
	"fmt"

	"atm-design/session"

	. "github.com/fatih/color"
)

func logout(...string) (Message, error) {
	s := session.GetSession()

	if !s.IsAuthorized() {
		return Message{YellowString("No account is currently authorized.")},
			fmt.Errorf("no authorized account")
	}

	acc := s.AccountID()
	s.Stop()
	return Message{GreenString("Account %d logged out.", acc)}, nil
}
