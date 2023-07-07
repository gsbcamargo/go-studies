package main

import "fmt"

func main() {
	var user1 user
	fmt.Printf("Struct variable without initialization: %v\n", user1) // each field is given it's default value (empty string and 0, in this scenario)
	user1.name = "Gabriel"
	fmt.Printf("User name: %s", user1.name)
}

type user struct {
	name string
	age  uint8
}
