package main 

import (
	"errors"
	"fmt"
)

type Logger struct{}

func (lgr *Logger) Log(message string) error {
	if len(message) == 0 {
		return errors.New("Message is Empty")
	}
	fmt.Println(message)

	return nil
}
