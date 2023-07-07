package main

import "fmt"

func main() {
	var user1 user
	fmt.Printf("Struct variable without initialization: %v\n", user1) // each field is given it's default value (empty string and 0, in this scenario)
	user1.name = "Gabriel"
	user1.age = 30
	fmt.Printf("User name: %s\n", user1.name)
	fmt.Printf("User age: %d", user1.age)

	birthPlace := birthPlace{"Brazil,", "Fictional Province,", "Fictional City"}

	user2 := user{"Lhara", 19, birthPlace} // brackets not parenthesis
	fmt.Printf("Second way of initalizing struct, user name: %s\n", user2.name)
	fmt.Printf("Second way of initalizing struct, user age: %d\n", user2.age)
	fmt.Printf("Second way of initalizing struct, user birthplace: %v\n", user2.birthPlace)

}

type user struct {
	name       string
	age        uint8
	birthPlace birthPlace
}

type birthPlace struct {
	country  string
	province string
	city     string
}
