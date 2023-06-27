package helper

import "fmt"

func Write() {
	fmt.Println("I'm inside a helper package.")
	privateWrite()
}

func privateWrite() {
	fmt.Println("Hi! I'm private! This function cannot be accessed outside of the 'helper' package.")
}
