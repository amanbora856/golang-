package main

import "fmt"

//for->only contructs for looping
func main() {

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i = 0; i < 10; i++ {
		fmt.Println((i))
	}
}
