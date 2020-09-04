package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Printf("Enter sentence:")
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.ToLower(strings.TrimSpace(line))
		if (line[0] == 'i' && line[len(line)-1] == 'n' && strings.Contains(line, "a")){
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}