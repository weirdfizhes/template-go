package logger

import (
	"fmt"
	"log"
	"os"
)

func LogPrintError(text string, err error) {
	fmt.Println("============================================")
	if text != "" {
		log.Printf("MESSAGE: %s\n", text)
	}
	log.Printf("ERROR: %s\n", err)
	fmt.Println("============================================")
}

func LogFatalError(text string, err error) {
	fmt.Println("============================================")
	if text != "" {
		log.Printf("MESSAGE: %s\n", text)
	}
	log.Printf("FATAL: %s\n", err)
	fmt.Println("============================================")
	os.Exit(1)
}

func LogPrintSuccess(text string, data interface{}) {
	fmt.Println("============================================")
	log.Printf("SUCCESS: %s\n", text)
	if data != nil {
		log.Printf("DATA: %s\n", data)
	}
	fmt.Println("============================================")
}
