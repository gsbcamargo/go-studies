package main

import (
	"fmt"
	"packages/helper"

	"github.com/badoux/checkmail"
	// go mod tidy -> to remove unused dependencies
)

func main() {
	fmt.Println("Hello, World!")
	helper.Write()

	correct := checkmail.ValidateFormat("gabriel@gabriel.dev")
	wrongFormat := checkmail.ValidateFormat("1234")
	fmt.Println(correct)
	fmt.Println(wrongFormat)
}
