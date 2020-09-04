package main
import (
	"fmt"
	"strings"
	"strconv"
	"sort"
  )

func main() {
	s := make([]int, 0, 3)
	for true {
		var input string
		fmt.Printf("Enter interger (enter X to exit): ")
		fmt.Scan(&input)
		i, err := strconv.Atoi(input)
		if(err == nil) {
			s = append(s, i)
			sort.Ints(s)
			fmt.Println(s)
		} else if(strings.ToLower(input) == "x") {
			break
		} else {
			fmt.Println("Invalid input!")
		}
	}
}
