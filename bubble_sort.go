package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// channle to arrange communication between go routines
var ch = make(chan []int)
var wg sync.WaitGroup

func random_generator(lists *[]int) {
	for i := 0; i < 100; i++ {
		*lists = append(*lists, rand.Intn(1000))
	}
	fmt.Printf("%v\n", *lists)
	ch <- *lists
	wg.Done()
}
func bubble() {
	if <-ch {
		fmt.Printf("%+v\n", "We can implement bubble sort now")
		fmt.Printf("%+v\n", "We got the list now")
		fmt.Printf("%+v\n", lists)
	}
	wg.Done()
}

func main() {
	var lists []int
	wg.Add(2)
	go random_generator(&lists)
	go bubble()
	fmt.Printf("%v\n", lists)
	wg.Wait()
}
