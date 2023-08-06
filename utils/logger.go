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

var Logger logger

/*
A log generator, it can give several options, and includes level choices.
*/
type logger struct {
	Level int
}

func PrintNewLog(msg string) {
	log.Print(msg)
}

func (l logger) Debug(msg string) {
	if l.Level < 1 {
		log.Println(DEBUG + msg)
	}
}

func (l logger) Info(msg string) {
	if l.Level < 2 {
		log.Println(INFO + msg)
	}
}

func (l logger) Warning(msg string) {
	if l.Level < 3 {
		log.Println(WARNING + msg)
	}
}

func (l logger) Error(msg string) {
	if l.Level < 4 {
		log.Panicln(ERROR + msg)
	}
}

func (l logger) Fatal(msg string) {
	log.Fatalln(FATAL + msg)
}
