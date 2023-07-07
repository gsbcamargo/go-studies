package main

import "fmt"

func main() {
	var user1 user
	user1.name = "Gabriel"
	fmt.Printf("User name: %s", user1.name)
}

type user struct {
	name string
	age  uint8
}
