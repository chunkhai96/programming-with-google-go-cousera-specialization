package main
import (
	"fmt"
  )

func main() {
	var input int
	data := make([]int, 0)

	for i:=0; i<10; i++ {
		fmt.Printf("%d Enter integer: ", i+1)
		fmt.Scan(&input)
		data = append(data, input)
	}

	BubbleSort(&data)
	fmt.Println("Sort result")
	fmt.Println(data)
}

func BubbleSort(data *[]int) {
	for i:=len(*data); i>0 ; i-- {
		for j:=0; j<i-1; j++ {
			if((*data)[j] > (*data)[j+1]) {
				Swap(&(*data)[j], &(*data)[j+1])
			}
		}
	}
}

func Swap(num1 *int, num2 *int) {
	var temp int
	temp = *num1
	*num1 = *num2
	*num2 = temp
}
