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

func main() {
	ch := make(chan bool)
	go foo(ch)
	<- ch
	go bar(ch)
	<- ch
}
