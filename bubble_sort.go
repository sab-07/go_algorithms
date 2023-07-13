package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"time"
)

// channel to arrange communication between go routines
var (
	ch      = make(chan *[]int)
	wg      sync.WaitGroup
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	orange  = "\033[38;5;208m"
	magenta = "\033[35m"
	blink   = "\033[5m"
	cyan    = "\033[36m"
)

func random_generator(lists *[]int) {
	// Generator generates 100 randow nums and append to list
	for i := 0; i < 10; i++ {
		*lists = append(*lists, rand.Intn(1000))
	}
	fmt.Printf(yellow+"Generated:\n%v\n"+reset, *lists)
	// fmt.Println("")
	// ch and lists are of same type. So no need to have * or &
	ch <- lists
	wg.Done()
}

// bubble func do the sorting

func bubble() {
	lists := <-ch
	// fmt.Printf("%T\n", lists)
	fmt.Printf(green+"Input to bubble sort:\n%v\n"+reset, *lists)
	count := 0
	lenOfList := len(*lists)
	for i := 0; i < lenOfList; i++ {
		for j := 0; j < lenOfList-1; j++ {
			count++
			perc := (float64(count)) / float64(lenOfList*(lenOfList-1)) * 100
			fmt.Printf("\033[6;1HProgress:%v", math.Floor(perc))
			// fmt.Println(orange, "Compare: ", (*lists)[j], (*lists)[j+1], reset)
			fmt.Printf("\033[7;12H%v", "    ")
			fmt.Printf(orange+"\033[7;1H\033[K[%v %v]"+reset, (*lists)[j], (*lists)[j+1])

			fmt.Printf(cyan+"\033[8;1HSoring:%v", *lists)
			if (*lists)[j] > (*lists)[j+1] {
				fmt.Printf(magenta + "\r\033[7;11H--> swap " + reset)
				// fmt.Println(magenta, "swap: ", (*lists)[j], "is bigger than", (*lists)[j+1], reset)
				(*lists)[j], (*lists)[j+1] = (*lists)[j+1], (*lists)[j]

				fmt.Printf(cyan+"\033[8;1HSoring:%v", *lists)
			}
			time.Sleep(time.Millisecond * 200)
		}
	}
	// fmt.Println(count)
	wg.Done()
}

// just call go routines and use wait groups for execution to be over

func main() {
	var lists []int
	wg.Add(2)
	cmd := exec.Command("clear")

	cmd.Stdout = os.Stdout
	cmd.Run()
	go random_generator(&lists)
	go bubble()
	wg.Wait()
	for i := 0; i < 1; i++ {
		// fmt.Printf(blink+"\rAfter the sort:\n%v\n"+reset, lists)
		fmt.Printf(blink+red+"\rSorted:%v"+reset, lists)
	}
	time.Sleep(time.Second * 5)
}
