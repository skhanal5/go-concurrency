package main

import (
	"fmt"
	"time"
)

/*
	Here we introduce a new construct: for-range loop over
	a channel. Observe how the output is deterministic.
*/
func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i<10; i++ {
			ch <- i
		}
		close(ch)
	}()
	
	go func() { 
		// goroutine pauses until a value is available
		// or the channel is closed
		for v := range ch {
			fmt.Println(v)
		}
	}()
	time.Sleep(100 * time.Millisecond)
}