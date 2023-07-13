package main

import (
	"fmt"
	"math/rand"
)

var ch = make(chan int)

// Functions checks max value return it into channel
func maxValue(slice []int, ch chan int) {
	max := slice[0]
	for i := 0; i < len(slice); i++ {
		if max < slice[i] {
			max = slice[i]
		}
	}
	ch <- max
}

func main() {
	var lists []int

	// Generating 100 random numbers
	for i := 0; i < 100; i++ {
		lists = append(lists, rand.Intn(1000))
	}

	fmt.Printf("Generated random sequence %v\n", lists)
	go maxValue(lists, ch)
	fmt.Printf("max is %v", <-ch)
}
