package main

import (
	"fmt"
	"time"
)

/*
 This is a slight adjustment to module 4. Instead of writing
 the value 6 immediately after, we first read the first 5 values
 written to the buffer.

 Run the code. Do you notice something off about the output?
*/
func main() {
	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		fmt.Printf("We have %d values in the buffer\n", len(ch))
	}()
	// These run non-deterministically, maybe some values were read
	go func() {
		for i := 0; i < 5; i++ {
			v := <- ch
			fmt.Printf("Reading value %d from the buffer\n", v)
		}
		fmt.Printf("We have %d values in the buffer \n", len(ch))
	}()
	time.Sleep(100 * time.Millisecond)
	ch <- 6 // What happens if we do a sixth write...
	fmt.Printf("The last value in the buffer: %d\n", (<-ch)) 
	fmt.Printf("We have %d values in the buffer \n", len(ch))
}