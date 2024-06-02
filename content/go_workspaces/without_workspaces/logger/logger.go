package logger

import "fmt"

func Info(msg string) {
	fmt.Printf("Info: %s\n", msg)
}

func Error(err error) {
	fmt.Printf("Error: %v\n", err)
}
