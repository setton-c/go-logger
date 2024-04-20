package main

import "strconv"

type Person struct {
	name string
	age  int
}

func (pers *Person) stringify() string {
	s := "Name: " + pers.name + " Age: " + strconv.Itoa(pers.age)

	return s
}
