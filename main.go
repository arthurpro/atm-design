package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"atm-design/command"
	"atm-design/db"
	"atm-design/log"
	"atm-design/session"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
)

var (
	whitespaces = regexp.MustCompile(`\s+`)
	suggestions []prompt.Suggest

	output chan []string
)

func main() {
	log.Init()
	defer log.Close()
	_, err := db.Connect("db.sqlite")
	if err != nil {
		log.Logger.Fatalln(err)
	}

	output = make(chan []string)
	session.SetTimeout(2 * time.Minute)
	session.SetOutput(&output)
	go printer()

	for n, d := range command.Suggestions() {
		suggestions = append(suggestions, prompt.Suggest{Text: n, Description: d})
	}

	output <- []string{color.BlueString("ATM Design Test. Enter your command:")}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionMaxSuggestion(0),
		prompt.OptionPrefix(""),
	)
	p.Run()
}

func executor(t string) {
	go session.GetSession().Ping()

	t = whitespaces.ReplaceAllString(t, " ")
	t = strings.Trim(t, " ")
	args := strings.Split(t, " ")

	log.Logger.Infoln(">", args)
	msg, err := command.Executor(args...)
	if err != nil {
		log.Logger.Errorln("<", err)
	}
	output <- msg
}

func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func printer() {
	for msg := range output {
		if len(msg) > 0 {
			fmt.Println(strings.Join(msg, "\n"))
		}
	}
}
