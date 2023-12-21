package utils

import (
	"log"
	// "os"
)

func InitLog(name string) error {
	// logFile, err := os.OpenFile(".log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	// log.SetOutput(logFile)
	return nil
}

func Log(msg string) error {
	log.Println(msg)
	// TODO: implmenet
	return nil
}
