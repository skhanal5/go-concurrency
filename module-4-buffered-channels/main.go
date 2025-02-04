package main

import (
	"fmt"
	"time"
)

/*
	This will cause a deadlock. Can you see why?
*/
func main() {
	// Here we have a buffered channel with size 5
	// This means we can do 5 non-blocking writes
	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		fmt.Printf("We have %d values in the buffer", len(ch))
	}()
	time.Sleep(100 * time.Millisecond)
	ch <- 6 // What happens if we do a sixth write...
}