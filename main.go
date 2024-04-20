package main

import (
	"fmt"
)

func main() {
	generator := Random_people_generator{}
	people := generator.generate_random_people(5)
	logr := Logger{}

	for i := 0; i < len(people); i++ {
		err := logr.Log(people[i].stringify())
		if err != nil {
			fmt.Println("[ERROR]: ", err)
			return
		}
	}

	err := logr.Log("")
	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

}
