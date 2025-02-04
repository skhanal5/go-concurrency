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
	Compare output with module 1.
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	wg.Wait()
	wg.Add(1)
	go bar(&wg)
	wg.Wait()
}
