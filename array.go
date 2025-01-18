package main

import "fmt"

func main() {

	var nums [5]int
	var vals [8]bool
	numss := [5]int{1, 2, 7, 8, 5}
	num2Daray := [2][2]int{{1, 2}, {1, 2}}
	fmt.Println(num2Daray)
	fmt.Println(numss)
	fmt.Println(vals)
	fmt.Println(len(nums))
	fmt.Println(nums) //default behaviour of array that 0 will be added;
}
