package logger

import (
	"log"
	"os"
	"sync"
	"time"
)

type Logger struct {
	mutex  sync.Mutex
	logger *log.Logger
}

func NewLogger() *Logger {
	date := time.Now().Format("2006-01-02")
	fileName := "logs_" + date + ".log"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file %v", err)
	}
	return &Logger{
		logger: log.New(file, "", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logger.Println("INFO:", msg)
}

func (l *Logger) Error(msg string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logger.Println("ERROR:", msg)
}
