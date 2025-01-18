package main

import "fmt"

const age = 30

func main() {
	name := 10
	const (
		port = 500
		host = "dev"
	)
	fmt.Println(host, port)
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(2 + 2)
	fmt.Println("code")
	fmt.Println(5 / 9)
}
