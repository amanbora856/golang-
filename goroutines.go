package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("I am not from goroutines")
	var wg sync.WaitGroup
	collection := []int{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, v := range collection {
		go func(val int) {
			wg.Add(1)
			print(val, &wg)
		}(v)
	}
	wg.Wait()
}

func print(input int, w *sync.WaitGroup) {
	fmt.Println("print %", input)
	defer w.Done()

}
