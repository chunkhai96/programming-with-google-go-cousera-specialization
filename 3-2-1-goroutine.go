package main
import (
	"fmt"
	"time"
)

func run(str string){
	for i:=0; i<100; i++ {
		fmt.Printf("%s: %d\n", str, i)
	}
}

func main() {
	
	// goroutine allows program to run concurrently
	go run("goroutine 1")
	go run("goroutine 2")

	time.Sleep(time.Second)
}