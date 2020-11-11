package main

import (
	"fmt"
)

func main() {

	arr := []int{8, 7, 9, 6, 4, 8, 0, 7, 4, 3}

	// traditional for loop
	fmt.Println("Traditional FOR loop")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("\nRange FOR loop with only value.")
	for _, arrElm := range arr {
		fmt.Print(arrElm, " ")
	}

	fmt.Println("\nRange FOR loop with index and value.")
	for index, arrElm := range arr {
		fmt.Println(index, " ", arrElm)
	}

	fmt.Println("\nFOR loop with condition.")
	i := 0
	for len(arr) != i {
		fmt.Println(i, " ", arr[i])
		i++
	}
}
