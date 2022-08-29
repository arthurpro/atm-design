package session

import (
	"fmt"
	"time"

	"atm-design/log"

	"github.com/fatih/color"
)

type ping bool
type stop bool

type Session struct {
	accountID *uint
	ticker    *time.Ticker
	ping      chan ping
	stop      chan stop
}

var timeout = 2 * time.Minute

var session *Session

var output *chan []string

func NewSession() (s *Session) {
	session = &Session{}
	session.ping = make(chan ping)
	session.stop = make(chan stop)
	return session
}

func GetSession() (s *Session) {
	if session == nil {
		session = NewSession()
	}
	return session
}

func SetTimeout(t time.Duration) {
	timeout = t
}

func SetOutput(o *chan []string) {
	output = o
}

func (s *Session) AccountID() uint {
	return *s.accountID
}
func (s *Session) IsAuthorized() bool {
	return s.accountID != nil
}

func (s *Session) Start(acc uint) {
	s.accountID = &acc
	s.ticker = time.NewTicker(timeout)

	for {
		select {
		case <-s.ticker.C:
			msg := []string{"", color.YellowString("Authorization expired after %s of inactivity.", timeout.String())}
			log.Logger.Errorln("session timeout")
			if output != nil {
				*output <- msg
			} else {
				fmt.Println(msg)
			}

			s.accountID = nil
			s.ticker.Stop()
			return
		case <-s.ping:
			if !s.IsAuthorized() {
				return
			}
			s.ticker.Reset(timeout)
		case <-s.stop:
			return
		}
	}
}

func (s *Session) Ping() {
	if s.IsAuthorized() {
		s.ping <- true
	}
}

func (s *Session) Stop() {
	if s.IsAuthorized() {
		s.accountID = nil
		s.stop <- true
	}
	if s.ticker != nil {
		s.ticker.Stop()
	}
}
