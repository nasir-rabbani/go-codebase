package main

import (
	"fmt"
	"sync"
)

func f1() {
	a := 100000000
	for a == 0 {
		a--
	}
	// fmt.Println("done f1")
}

func f2() {
	a := 100000000
	for {
		if a == 0 {
			break
		} else {
			a--
		}
	}
	// fmt.Println("done f2")
}

func main() {
	fmt.Println("Hello, playground")

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		f2()
		fmt.Println("F2 called")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		f1()
		fmt.Println("F1 called")
		wg.Done()
	}()

	wg.Wait()

}
