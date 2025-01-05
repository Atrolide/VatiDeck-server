// Package logger provides a custom logger for the VatiDeck server.
package logger

import (
	"log"
	"os"
)

// Extends the standard log.Logger.
type Logger struct {
	*log.Logger
}

// InitLogger configures and returns a logger instance.
func InitLogger() *Logger {
	// Set up logger with timestamp and file/line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
	return &Logger{log.Default()}
}

// Info logs an informational message.
func (l *Logger) Info(msg string) {
	l.Printf("INFO: %s", msg)
}

// Error logs an error message.
func (l *Logger) Error(msg string) {
	l.Printf("ERROR: %s", msg)
}
