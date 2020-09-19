package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct { a string}

func (c *Cow) Eat() {
	fmt.Println("grass")
}

func (c *Cow) Move() {
	fmt.Println("walk")
}

func (c *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct { a string}

func (b *Bird) Eat() {
	fmt.Println("worms")
}

func (b *Bird) Move() {
	fmt.Println("fly")
}

func (b *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct { a string}

func (s *Snake) Eat() {
	fmt.Println("mice")
}

func (s *Snake) Move() {
	fmt.Println("slither")
}

func (s *Snake) Speak() {
	fmt.Println("hss")
}

func main() {
	dict :=  make(map[string]Animal)

	for {
		fmt.Printf(">")
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.ToLower(strings.TrimSpace(line))
		data := strings.Split(line, " ")

		if data[0] == "newanimal" {
			switch data[2] {
			case "cow":
				var cow Animal = &Cow{}
				dict[data[1]] = cow
			case "bird":
				var bird Animal = &Bird{}
				dict[data[1]] = bird
			case "snake":
				var snake Animal = &Snake{}
				dict[data[1]] = snake
			default:
				fmt.Println("No such animal!")
			}

			fmt.Println("Created it!")

		} else if data[0] == "query" {
			if _, ok := dict[data[1]]; ok {
				switch data[2] {
				case "eat":
					dict[data[1]].Eat()		
				case "move":
					dict[data[1]].Move()
				case "speak":
					dict[data[1]].Speak()
				default:
					fmt.Println("No such move!")
				}
			} else {
				fmt.Println("No animal named", data[1])
			}
		} 
	}
}