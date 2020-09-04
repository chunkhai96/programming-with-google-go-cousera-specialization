package main
import "fmt"

func main() {
	for true {
		var input float64
		fmt.Printf("Please enter number: ")
		fmt.Scan(&input)
		fmt.Printf("%d\n", int(input))
	}
}
