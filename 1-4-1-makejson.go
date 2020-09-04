package main
import (
	"fmt"
	"encoding/json"
  )

func main() {
	data :=  make(map[string]string)
	var input string
	fmt.Println("Enter Name: ")
	fmt.Scan(&input)
	data["name"] = input
	fmt.Println("Enter Address: ")
	fmt.Scan(&input)
	data["address"] = input

	jsonString, _ := json.Marshal(data)
    fmt.Println(string(jsonString))
}
