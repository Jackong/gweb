/**
 * User: jackong
 * Date: 10/28/13
 * Time: 10:58 AM
 */
package log

import (
	"os"
	"fmt"
	"log"
	"time"
)

type Level string

const (
	DEBUG Level = "D"
	INFO Level = "I"
	WARNING Level = "W"
	ERROR Level = "E"
)

var (
	dir string
	logger *log.Logger
	logFile *os.File
)

func init() {
	dir = "log"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println("warning:", err)
		os.Exit(1)
	}
}

func getLog() *log.Logger {
	fileName := dir + "/" + time.Now().Format("2006-01-02") + ".log"
	if logFile != nil {
		if logFile.Name() == fileName {
			return logger
		}
		logFile.Sync()
		logFile.Close()
		logFile = nil
	}

	var err error = nil
	logFile, err = os.OpenFile(fileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("warning:", err)
	}
	logger = log.New(logFile, "", log.Ldate | log.Ltime | log.Lshortfile)
	return logger
}

func Debug(v ... interface {}) {
	Print(DEBUG, v)
}

func Info(v ... interface {}) {
	Print(INFO, v)
}

func Warning(v ... interface {}) {
	Print(WARNING, v)
}

func Error(v ... interface {}) {
	Print(ERROR, v)
}

func Print(level Level, v ... interface {}) {
	Output(0, level, v)
}

func Output(depth int, level Level, v ... interface {}) {
	getLog().Output(4 + depth, fmt.Sprintf("[%v] %v", level, v))
}

func Close() {
	logFile.Close()
}
