package constants

import (
	"bytes"
	"log"
	"os"
	"time"
)

const (
	VERSION     string = "b0.1"
	MODULE_NAME string = "KONN"
)

const (
	INFO    = "[INFO]"
	DEBUG   = "[DEBUG]"
	WARNING = "[WARNING]"
	ERROR   = "[ERROR]"
	FATAL   = "[FATAL]"
)

var Host string = "xxx"
var LogEvent chan string

// Generate a log to your target string.
func LogGenerate(msg string, logLevel string) string {
	if msg == FATAL {
		defer os.Exit(1)
		defer time.Sleep(5 * 1e9)
	}
	buf := new(bytes.Buffer)
	logger := log.New(buf, logLevel, log.Ldate|log.Ltime)
	logger.Println(msg)
	return buf.String()
}
