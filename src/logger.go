package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var Logger *log.Logger

func InitLogger(logDir string) error {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	file := fmt.Sprintf("%s/run-%s.log",
		logDir,
		time.Now().Format("20060102-150405"),
	)

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	Logger = log.New(f, "", log.LstdFlags)
	return nil
}

func Info(msg string) {
	fmt.Printf("[INFO] %s\n", msg)
	if Logger != nil {
		Logger.Println("[INFO]", msg)
	}
}

func Warn(msg string) {
	fmt.Printf("[WARN] %s\n", msg)
	if Logger != nil {
		Logger.Println("[WARN]", msg)
	}
}

func Error(msg string) {
	fmt.Printf("[ERROR] %s\n", msg)
	if Logger != nil {
		Logger.Println("[ERROR]", msg)
	}
}
