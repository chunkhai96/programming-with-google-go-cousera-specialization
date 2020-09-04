package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
  )

func main() {
	type Name struct {
		fname string
		lname string
	}
	var filename string
	names := make([]Name, 0)

	fmt.Printf("Enter filename: ")
	fmt.Scan(&filename)
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		n := Name{input[0], input[1]}
		names = append(names, n)
	}
	
	for _, name := range(names) {
		fmt.Printf("First Name : %s \t Last Name: %s\n", name.fname, name.lname)
	}
}
