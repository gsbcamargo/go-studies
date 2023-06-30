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

	variable := "variable"

	const constant string = "constant"

	correct := checkmail.ValidateFormat("gabriel@gabriel.dev") // if correct returns nil
	wrongFormat := checkmail.ValidateFormat("1234")
	printedVariable := fmt.Sprintf("I'm a variable! Look: %s.", variable) // string interpolation
	printedConstant := fmt.Sprintf("I'm a constant! Look: %s.", constant)
	fmt.Println(printedVariable)
	fmt.Println(printedConstant)
	fmt.Println(correct)
	fmt.Println(wrongFormat)
}
