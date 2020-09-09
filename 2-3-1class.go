package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	
	for {
		var animal Animal

		fmt.Printf(">")
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.ToLower(strings.TrimSpace(line))
		data := strings.Split(line, " ")

		switch data[0] {
		case "cow":
			animal = Animal {"grass", "walk", "moo"}
		case "bird":
			animal = Animal {"worms", "fly", "peep"}
		case "snake":
			animal = Animal {"mice", "slither", "hss"}
		}

		switch data[1] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()	
		}
	}
}