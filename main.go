package main

import (
	"github.com/setton-c/go-logger/pkg/logger"
)

func main() {
	logr := logger.Create(logger.INFO)
	generator := Random_people_generator{}
	
	people, err := generator.Generate_random_people(0)
	if err != nil {
		logr.Log(logger.ERROR, err.Error())
	}

	for i := 0; i < len(people); i++ {
		logr.Log(logger.INFO, people[i].stringify())
	}
	
	
}
