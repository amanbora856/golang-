package main

import "fmt"

// Safe division function using recover()
func divide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("Division by zero is not allowed!")
	}
	return a / b
}

func main() {
	fmt.Println("10 / 2 =", divide(10, 2)) // Output: 5
	fmt.Println("5 / 0 =", divide(5, 0))   // Handles panic and recovers
	fmt.Println("8 / 4 =", divide(8, 4))   // Output: 2
}
