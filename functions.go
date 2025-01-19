package main

import "fmt"

func main() {
	final := add(5, 6)
	fmt.Println("sum is ", final)
	output1, output2, output3 := getLanguages()
	fmt.Println("output is ", output1, output2, output3)
	totalsum := calc(5, 6, 7, 8, 9)
	fmt.Println(totalsum)
}

func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, int) {
	return "aman", "bora", 5
}

func calc(nums ...int) int {
	total := 0
	for _, val := range nums {
		total = total + val
	}

	return total
}
