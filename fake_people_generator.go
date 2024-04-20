package main

import (
	"math/rand"
)

type Random_people_generator struct{}

func (ng *Random_people_generator) generate_random_people(amount int) []Person{
	var possible_names = []string{"Bob", "Vakho", "Misho", "Nika", "Giorgi"}
	var people []Person

	for i := 0; i < amount; i++ {
		random_index := rand.Intn(len(possible_names))
		name := possible_names[random_index]
		age := rand.Intn(50) + 20

		pers := Person{
			name: name,
			age:  age,
		}

		people = append(people, pers)
	}

	return people
}
