package main

import (
	"io"
	"log"
	"os"
	"time"
)

type AccessMessage struct {
	Time   time.Time
	From   interface{}
	To     interface{}
	Status string
	Reason interface{}
}

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("open error.log fail: ", err)
	}

	Info = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)

}
