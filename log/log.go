package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger
var f *os.File

func Init() {
	Logger = logrus.New()
	var err error
	f, err = os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		Logger.Fatalf("error opening file: %v", err)
	}
	Logger.SetOutput(f)
}

func Close() {
	err := f.Close()
	if err != nil {
		Logger.Fatalf("error closing file: %v", err)
	}
}
