package main

import (
	"fmt"
)

func count(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go count(ch)
	fmt.Println("i was here", <-ch, <-ch)
	for num := range ch {
		fmt.Println(num)
	}
}
