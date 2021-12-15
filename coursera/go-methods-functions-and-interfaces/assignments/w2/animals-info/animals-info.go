package main

import (
	"fmt"
)

// TYPES::

type Animal struct {
	food, locomotion, noise string
}

type ChosenAnimal struct {
	animalType, functionality string
}

// METHODS::

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	snake := Animal{"mic", "slither", "hsss"}
	cow := Animal{"grass", "walk", "mooo"}
	bird := Animal{"worms", "fly", "peep"}

	//animals:= []string{"snake", "cow", "bird"}
	//functionalities:= []string{"move", "speak", "eat"}

	animal := ChosenAnimal{}
	for {
		fmt.Println("You either enter <snake>, <cow> or <bird>")
		fmt.Println("Available functionality: <eat>, <move> or <speak>")

		// get the animal from the user:
		fmt.Println()
		_, err := fmt.Scanf("%s %s", &animal.animalType, &animal.functionality)
		if err != nil {
			fmt.Println("Please choose functionality of the above ^^")
		}

		switch animal.animalType {
		case "snake":
			if animal.functionality == "eat" {
				fmt.Print("Snakes eat: ")
				snake.Eat()
			}
			if animal.functionality == "speak" {
				fmt.Print("Snakes speak: ")
				snake.Speak()
			}
			if animal.functionality == "move" {
				fmt.Print("Snakes move: ")
				snake.Move()
			}

		case "cow":
			if animal.functionality == "eat" {
				fmt.Print("Cows eat: ")
				cow.Eat()
			}
			if animal.functionality == "speak" {
				fmt.Print("Cows speak: ")
				cow.Speak()
			}
			if animal.functionality == "move" {
				fmt.Print("Cows move: ")
				cow.Move()
			}

		case "bird":
			if animal.functionality == "eat" {
				fmt.Print("Birds eat: ")
				bird.Eat()
			}
			if animal.functionality == "speak" {
				fmt.Print("Birds speak: ")
				bird.Speak()
			}
			if animal.functionality == "move" {
				fmt.Print("Birds move: ")
				bird.Move()
			}

		default:
			fmt.Println("Please choose animal of the above :)")
		}

		fmt.Println("————————————————————————")

		// reset animal struct
		animal.animalType = ""
		animal.functionality = ""
	}
}
