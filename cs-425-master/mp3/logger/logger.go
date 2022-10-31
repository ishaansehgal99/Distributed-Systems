package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	// InfoLogger : logging for info messages
	InfoLogger *log.Logger

	// WarningLogger : logging for warning messages
	WarningLogger *log.Logger

	// ErrorLogger : logging for error messages
	ErrorLogger *log.Logger
)

// Initiate the logger
func init() {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Print in the console and in the log as level INFO
func PrintInfo(args ...interface{}) {
	InfoLogger.Println("INFO: ", args)
	fmt.Println(args...)
}

// Print in the console and in the log as level WARN
func PrintWarning(args ...interface{}) {
	WarningLogger.Println("WARNING: ", args)
	fmt.Println(args...)
}

// Print in the console and in the log as level ERROR
func PrintError(args ...interface{}) {
	ErrorLogger.Println("ERROR: ", args)
	fmt.Println(args...)
}
