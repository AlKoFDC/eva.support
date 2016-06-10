package logger

import (
	"log"
	"os"
)

// Standard is the standard logger, that prints to os.Stdout.
var Standard = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

// Error is the error logger, that prints to os.Stderr
var Error = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
