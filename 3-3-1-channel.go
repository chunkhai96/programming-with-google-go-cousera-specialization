package main
import (
	"fmt"
	"sort"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func sorting(num int, arr []int, c chan<- []int) {
	sort.Ints(arr)
	fmt.Printf("Splited sorted array %d: %v\n", num, arr)
	c <- arr
}

func main() {

	partNum := 4
	arr := make([]int, 0)

	for {
		fmt.Printf("Enter integer (enter x to exit input): ")
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.ToLower(strings.TrimSpace(line))

		if line == "x" {
			break
		}

		data, err := strconv.Atoi(line)
		if err == nil {
			arr = append(arr, data)
		}
	}

	c := make(chan []int, partNum)
	remain := len(arr) % partNum
	startPos := 0

	for i:= 0; i<partNum; i++ {
		endPos := startPos + (len(arr) / partNum)
		if remain > 0 {
			endPos++
			remain--
		}
		go sorting(i, arr[startPos:endPos], c)
		startPos = endPos
	}

	result := make([]int, 0)

	for i:=0; i<partNum; i++ {
		result = append(result, <-c... )
	}

	fmt.Printf("Merged: %v\n", result)
	sort.Ints(result)
	fmt.Printf("Merged and sorted: %v\n",result)
}