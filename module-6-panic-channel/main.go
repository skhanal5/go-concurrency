package main

import (
	"time"
)

/*
	This time we write to a channel and try to close it.
	Then we write to it again. Observe what happens
*/ 
func main() {
	ch := make(chan int)
	
	go func() {
		for i := 0 ; i<5; i++ {
			ch <- i
		}
		close(ch)
	}()
	time.Sleep(100 * time.Millisecond)
	ch <- 10
}