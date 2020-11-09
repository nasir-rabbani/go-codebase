package main

import (
	"fmt"
	"sync"
	"time"
)

func printName(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, name)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("Done printName")
}

func printMsg(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("Done printMsg")
}

func main() {
	fmt.Println("concurrency")

	var wg sync.WaitGroup

	wg.Add(1)
	go func(name string) {
		printName(name)
		wg.Done()
	}("Nasir")

	wg.Add(1)
	go func(msg string) {
		printMsg(msg)
		wg.Done()
	}("Hi, I am Nasir.")

	wg.Wait()
}
