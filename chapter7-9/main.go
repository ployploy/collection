package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	//key start 0
	for i, number := range slice {
		fmt.Println(i, number)
	}
}
