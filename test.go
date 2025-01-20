package main

import "fmt"

func greet(name string, greetings ...string) {
	if len(greetings) > 0 {
		fmt.Printf("%s, %s, %s!\n", greetings[0], greetings[1], name)
	} else {
		fmt.Printf("Hello, %s!\n", name)
	}
}

func main() {
	greet("Alice")                         // Output: Hello, Alice!
	greet("Bob", "Good Morning", "dsfsdf") // Output: Good Morning, Bob!
}
