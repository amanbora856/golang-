package main

import "fmt"

// Define the calculator interface
type Calculator interface {
	Calculate(a int, b int) int
}

// Define the Adder struct implementing Calculator
type Adder struct{}

func (a Adder) Calculate(x int, y int) int {
	return x + y
}

// Define the Subtracter struct implementing Calculator
type Subtracter struct{}

func (s Subtracter) Calculate(x int, y int) int {
	return x - y
}

// Function to print results using any Calculator implementation
func PrintResult(c Calculator, a, b int) {
	fmt.Println("Result:", c.Calculate(a, b))
}

func main() {
	// Instantiate Adder and Subtracter
	add := Adder{}
	sub := Subtracter{}

	// Print results using both implementations
	PrintResult(add, 5, 6) // Result: 11
	PrintResult(sub, 5, 6) // Result: -1
}
