package main

import (
	"log"
	"os"
)

func main() {
	log.Println("This is a log message.")

	log.SetPrefix("INFO: ")
	log.Println("This is an info message.")

	// Log Flags
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("This is a log message with only date.")

	infoLogger.Println("This is an info message.")
	warnLogger.Println("This is an warning message.")
	errorLogger.Println("This is an error message.")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: ", err)
	}
	defer file.Close()

	infoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger := log.New(file, "WARN: ", log.Ldate|log.Ltime)
	errorLogger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message.")
	warnLogger.Println("This is a warning message")
	infoLogger.Println("This is a info message")
	errorLogger.Println("This is a error message")
}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
