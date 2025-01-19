package main

import "fmt"

func main() {

	fmt.Println("Slices Practice")
	m := make(map[string]int)
	m["rollno"] = 5
	m["class"] = 8
	fmt.Println(m["rollno	"])
	no := map[string]int{"phone": 1, "room": 1, "land": 1}
	values, truth := no["phone"]
	if truth {
		fmt.Println("value exist", values)
	} else {
		fmt.Println("not okay", "empty ")
	}
}
