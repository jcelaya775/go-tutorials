package main

import (
	"fmt"
	"time"
)

// Goroutines
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Channels
func sum(s []int, c chan int) {
	sum := 0
	for _, val := range s {
		sum += val
	}
	c <- sum // send sum to c
}

// Buffered Channels
func bufferedChannelsExample() {
	ch := make(chan int, 2)
	// Able to send multiple values before channel blocks
	ch <- 1
	ch <- 2
	// Able to receive until there is nothing left to receive and channel blocks
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Range and close
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// Select
func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// NOTE: Select will choose a random case that is ready to execute
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// Default Selection
func defaultSelectionExample() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)
	for {
		select {
		case currentTime := <-tick:
			fmt.Println("Tick.")
			fmt.Println(currentTime.Second())
		case endTime := <-boom:
			fmt.Println("BOOM!")
			fmt.Println(endTime.Second())
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
