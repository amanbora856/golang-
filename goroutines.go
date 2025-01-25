package main

import (
	"fmt"
	"sync"
)

func hello(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < num; i++ {
		fmt.Println("Output is", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(5, &wg)
	wg.Wait()
}
