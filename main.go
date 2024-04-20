package main

import (
	"github.com/setton-c/go-logger/pkg/logger"
)

func main() {
	generator := Random_people_generator{}
	people := generator.generate_random_people(5)
	logr := logger.Create(logger.INFO)

	for i := 0; i < len(people); i++ {
		logr.Info(people[i].stringify())
	}

	err := logr.Log("")
	if err != nil {
		logr.Error(err.Error())
		return
	}

}
