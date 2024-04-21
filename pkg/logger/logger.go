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

func (lgr Logger) Log(level Log_level, args ...interface{}) {
	level_str := strings.ToUpper(level.String())
	
	fmt.Print("["+level_str+"]: ")
	fmt.Println(args...)
}

func (lgr Logger) Logf(level Log_level, format_string string, args ...interface{}) {	
	level_str := strings.ToUpper(level.String())
	fmt.Printf("["+level_str+"]: " + format_string + "\n", args...)
}

// func (lgr *Logger) Info(msg string) {
// 	if lgr.log_level <= INFO {
// 		lgr.Log("[INFO] " + msg)
// 	}
// }
//
// func (lgr *Logger) Warning(msg string) {
// 	if lgr.log_level <= WARNING {
// 		lgr.Log("[WARNING] " + msg)
// 	}
// }
//
// func (lgr *Logger) Error(msg string) {
// 	if lgr.log_level <= ERROR {
// 		lgr.Log("[ERROR] " + msg)
// 	}
// }
