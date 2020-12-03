package main

import "fmt"

func main() {
	// s := []int{20, 1990, 12, 1110, 1, 59, 12, 15, 120, 1110, -1, -19, 3000}
	s := []int{-2, 20, -3}
	max := s[0]
	var nextToMax int

	for _, val := range s {
		if val > max {
			nextToMax = max
			max = val
		} else if val > nextToMax {
			nextToMax = val
		} else if max == nextToMax {
			nextToMax = val
		}
	}
	fmt.Println(max)
	fmt.Println(nextToMax)
}
