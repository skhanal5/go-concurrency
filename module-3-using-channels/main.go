package main

import (
	"fmt"
)

func foo(ch chan bool) {
	fmt.Println("In function: foo()")
	ch <- true
}

func bar(ch chan bool) {
	fmt.Println("In function: bar()")
	ch <- true
}
/*
	Same as module 2. Why?
*/
func main() {
	// Using unbuffered channel.
	// For each write to this channel, each writing
	// goroutine pauses until another goroutine reads 
	// Same behavior for writes.
	// Thus, for a read/write to happen -> we need two concurrently running goroutines
	ch := make(chan bool) 
	go foo(ch)
	<- ch
	go bar(ch)
	<- ch
}
