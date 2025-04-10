package utils

import (
	"log"
	"os"
)

// InitLogger 初始化日志记录器
func InitLogger() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetPrefix("[Blog System] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
