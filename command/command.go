package command

import (
	"fmt"

	"atm-design/session"

	. "github.com/fatih/color"
)

var commands map[string]command

type command struct {
	arguments    string
	argsNumber   int
	authRequired bool
	handler      handler
}

type Message []string

type handler func(args ...string) (Message, error)

func init() {
	commands = map[string]command{
		"authorize": {
			"<account_id> <pin>",
			2,
			false,
			authorize,
		},
		"balance": {
			"",
			0,
			true,
			balance,
		},
		"deposit": {
			"<amount>",
			1,
			true,
			deposit,
		},
		"withdraw": {
			"<amount>",
			1,
			true,
			withdraw,
		},
		"history": {
			"",
			0,
			true,
			history,
		},
		"logout": {
			"",
			0,
			false,
			logout,
		},
		"end": {
			"",
			0,
			false,
			end,
		},
	}
}

func Executor(args ...string) (Message, error) {
	c := args[0]
	args = args[1:]

	if len(c) == 0 {
		return Message{""}, nil
	}

	if cmd, ok := commands[c]; !ok {
		return Message{RedString("Unknown command '%s'.", c)},
			fmt.Errorf("unknown command '%s'", c)
	} else {
		s := session.GetSession()
		if cmd.authRequired && !s.IsAuthorized() {
			return Message{YellowString("Authorization required.")},
				fmt.Errorf("authorization required")
		}

		if cmd.argsNumber != len(args) {
			return Message{RedString("Wrong number of arguments for '%s', requires %d argument(s).", c, cmd.argsNumber)},
				fmt.Errorf("wrong number of arguments")
		}

		msg, err := (cmd.handler)(args...)
		return msg, err
	}
}

func Suggestions() map[string]string {
	s := make(map[string]string)
	for name, c := range commands {
		s[name] = c.arguments
	}
	return s
}
