package main
import (
	"fmt"
	"time"
)

func run(str string){
	for i:=0; i<5; i++ {
		num += 2
		fmt.Printf("%s: %d\n", str, num)
	}
}

func run2(str string){
	for i:=0; i<5; i++ {
		num += 5
		fmt.Printf("%s: %d\n", str, num)
	}
}

var num int

func main() {
	
	/*
	** - race is a condition where two or more goroutine access to same resource
	** - since the goroutine execution is uncertain, this make the output difficult to predict
	** - it might cause bug in the application
	*/

	go run("goroutine 1")
	go run("goroutine 2")

	time.Sleep(time.Second)
}