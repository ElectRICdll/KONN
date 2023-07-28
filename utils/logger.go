package utils

import (
	"log"
)

const (
	DEBUG   = "[DEBUG]"
	INFO    = "[INFO]"
	WARNING = "[WARNING]"
	ERROR   = "[ERROR]"
	FATAL   = "[FATAL]"
)

/*
A log generator, it can give several options, and includes level choices.
*/
type Logger struct {
	Level int
}

func PrintNewLog(msg string) {
	log.Print(msg)
}

func (l Logger) Debug(msg string) {
	if l.Level < 1 {
		log.Println(DEBUG + msg)
	}
}

func (l Logger) Info(msg string) {
	if l.Level < 2 {
		log.Println(INFO + msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.Level < 3 {
		log.Println(WARNING + msg)
	}
}

func (l Logger) Error(msg string) {
	if l.Level < 4 {
		log.Panicln(ERROR + msg)
	}
}

func (l Logger) Fatal(msg string) {
	log.Fatalln(FATAL + msg)
}
