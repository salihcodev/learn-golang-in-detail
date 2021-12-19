package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
	GetName() string
}

func (c cow) GetName() string {
	return c.name
}

func (s snake) GetName() string {
	return s.name
}

func (b bird) GetName() string {
	return b.name
}

func (c cow) Eat() {
	fmt.Printf("%s eats grass\n", c.name)
}

func (c cow) Move() {
	fmt.Printf("%s walks\n", c.name)
}

func (c cow) Speak() {
	fmt.Printf("%s moos\n", c.name)
}

func (s snake) Eat() {
	fmt.Printf("%s eats mice\n", s.name)
}

func (s snake) Move() {
	fmt.Printf("%s slithers\n", s.name)
}

func (s snake) Speak() {
	fmt.Printf("%s hisses\n", s.name)
}

func (b bird) Eat() {
	fmt.Printf("%s eats worms\n", b.name)
}

func (b bird) Move() {
	fmt.Printf("%s flys\n", b.name)
}

func (b bird) Speak() {
	fmt.Printf("%s peeps and chirps\n", b.name)
}

type cow struct{ name string }
type snake struct{ name string }
type bird struct{ name string }

func main() {
	var animals []Animal
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("-____INSTRUCTIONS____-")
	fmt.Println("newanimal <name of animal> <type of animal>")
	fmt.Println("query <name of animal> <action of animal>")

	for {
		fmt.Print(">")
		userQuery, _, err := inputReader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		query := strings.Split(string(userQuery), " ")

		if len(query) != 3 {
			fmt.Println("Not enough parameters :)")
		}

		switch query[0] {
		case "newanimal":
			if query[2] == "cow" {
				animals = append(animals, cow{name: query[1]})
				fmt.Println("Created it!")
			} else if query[2] == "snake" {
				animals = append(animals, snake{name: query[1]})
				fmt.Println("Created it!")
			} else if query[2] == "bird" {
				animals = append(animals, bird{name: query[1]})
				fmt.Println("Created it!")
			}

		case "query":
			for _, animal := range animals {
				if animal.GetName() == query[1] {
					if query[2] == "move" {
						animal.Move()
					} else if query[2] == "eat" {
						animal.Eat()
					} else if query[2] == "speak" {
						animal.Speak()
					}
				}
			}
		}

	}

}
