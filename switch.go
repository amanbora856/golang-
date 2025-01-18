package main

import "fmt"

//for->only contructs for looping
func main() {
	i := 5
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("bestcolor")
	}
}
