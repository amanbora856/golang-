package main

import "fmt"

func main() {
	// Create a closure with an initial count of 0
	closure := add()

	// Call the closure with arguments and print results
	fmt.Println(closure(3)) // Output: 3 (0 + 3)
	fmt.Println(closure(2)) // Output: 5 (3 + 2)
	fmt.Println(closure(7)) // Output: 12 (5 + 7)
}

// add returns a closure that maintains state
func add() func(a int) int {
	count := 0 // captured variable, starts at 0
	return func(a int) int {
		count += a // modify the captured variable
		return count
	}
}
