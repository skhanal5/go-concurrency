package main

import (
	"fmt"
	"sync"
)

func bar(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("In function: bar()")
}

func foo(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("In function: foo()")
}

/*
	Run this code many times and make note
	of the output order.
*/
func main() {
	// wg was added because without it, the main function will exist immediately
	// before the goroutines have a chance to finish
	var wg sync.WaitGroup
	wg.Add(2)

	go foo(&wg)
	go bar(&wg)

	wg.Wait()
}
