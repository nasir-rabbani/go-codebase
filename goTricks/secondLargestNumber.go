package main

import "fmt"

func main() {
	s := []int{20, 1990, 12, 1110, 1, 59, 12, 15, 120, 1110, -1, -19, 3000}

	max := s[0]
	nextToMax := s[0]

	for _, val := range s {
		if val > max {
			max = val
		} else if val > nextToMax {
			nextToMax = max
		}
	}
	fmt.Println(max)
	fmt.Println(nextToMax)
}
