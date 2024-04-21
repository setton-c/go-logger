package logger

import (
	"fmt"
	"strings"
)

type Logger struct {
	log_level Log_level
}

func Create(log_level Log_level) *Logger {
	return &Logger{log_level: log_level}
}

func (lgr *Logger) Set_log_level(log_level Log_level) {
	lgr.log_level = log_level
}

func (lgr Logger) Log(args ...interface{}) {
	level_str := strings.ToUpper(lgr.log_level.String())
	
	fmt.Print("["+level_str+"]: ")
	fmt.Println(args...)
}

func (lgr Logger) Logf(format_string string, args ...interface{}) {	
	level_str := strings.ToUpper(lgr.log_level.String())
	fmt.Printf("["+level_str+"]: " + format_string + "\n", args...)
}

func (lgr *Logger) Info(args ...interface{}) {
	if lgr.log_level <= INFO {
		lgr.Log(args...)
	}
}

func (lgr *Logger) Infof(format_string string, args ...interface{}) {
	if lgr.log_level <= INFO {
		lgr.Logf(format_string, args...)
	}
}

func (lgr *Logger) Warning(args ...interface{}) {
	if lgr.log_level <= WARNING {
		lgr.Log(args...)
	}
}

func (lgr *Logger) Warningf(format_string string, args ...interface{}) {
	if lgr.log_level <= WARNING {
		lgr.Logf(format_string, args...)
	}
}

func (lgr *Logger) Error(args ...interface{}) {
	if lgr.log_level <= ERROR {
		lgr.Log(args...)
	}
}

func (lgr *Logger) Errorf(format_string string, args ...interface{}) {
	if lgr.log_level <= ERROR {
		lgr.Logf(format_string, args...)
	}
}
