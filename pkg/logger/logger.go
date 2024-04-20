package logger

import (
	"errors"
	"fmt"
)

type Log_level int

const (
	INFO Log_level = iota
	WARNING
	ERROR
)

type Logger struct{
	log_level Log_level
}

func Create(log_level Log_level) *Logger {
	return &Logger{log_level: log_level}
}

func (lgr *Logger) Set_log_level(log_level Log_level) {
	lgr.log_level = log_level
}

func (lgr *Logger) Log(message string) error {
	if len(message) == 0 {
		return errors.New("Message is Empty")
	}
	fmt.Println(message)

	return nil
}

func (lgr *Logger) Info(msg string) {
	if lgr.log_level <= INFO {
		err := lgr.Log("[INFO] " + msg)
		if err != nil {
			fmt.Println("[ERROR]: ", err)
			return
		}
	}
}

func (lgr *Logger) Warning(msg string) {
	if lgr.log_level <= WARNING {
		lgr.Log("[WARNING] " + msg)
	}
}

func (lgr *Logger) Error(msg string) {
	if lgr.log_level <= ERROR {
		lgr.Log("[ERROR] " + msg)
	}
}
