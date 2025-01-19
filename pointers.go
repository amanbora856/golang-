package main

import "fmt"

func main() {

	num := 8
	fmt.Println(&num)
	change(&num)
}

func change(num *int) {
	*num = 5
	fmt.Println(*num)
	fmt.Println(num)
}
