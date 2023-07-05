package main

import "fmt"

func main() {

	square := func(firstNumber int8) int8 {
		return firstNumber * firstNumber
	}

	sumOperation := sum(1, 2)
	concatenation := concat("Hello ", "World")
	concatenationFmt := concatFmt("Hello", "Universe")
	squared := square(4)
	sumOfDifferentSizes1, sumOfDifferentSizes2, sumOfDifferentSizes3, sumOfDifferentSizes4 := sumOfDifferentSizes(0, 9223372036854775807)
	fmt.Printf("sum: %d\n", sumOperation)
	fmt.Printf("concat: %s\n", concatenation)
	fmt.Printf("concatFmt: %s\n", concatenationFmt)
	fmt.Printf("inner func declared in variable: %d\n", squared)
	fmt.Printf("sumOfDifferentSizes1 (int8): %d\n", sumOfDifferentSizes1)
	fmt.Printf("sumOfDifferentSizes2 (int16): %d\n", sumOfDifferentSizes2)
	fmt.Printf("sumOfDifferentSizes3 (int32): %d\n", sumOfDifferentSizes3)
	fmt.Printf("sumOfDifferentSizes4 (int64): %d\n", sumOfDifferentSizes4)
}

func sum(firstNumber int8, secondNumber int8) int8 {
	return firstNumber + secondNumber
}

func concat(str1 string, str2 string) string {
	return str1 + str2
}

func concatFmt(str1 string, str2 string) string {
	concatenated := fmt.Sprintf("%s %s", str1, str2)
	return concatenated
}

func sumOfDifferentSizes(firstNumber, secondNumber int64) (int8, int16, int32, int64) {
	var sum8 int8 = 0 + 127
	var sum16 int16 = 0 + 32767
	var sum32 int32 = 0 + 2147483647
	var sum64 int64 = 0 + 9223372036854775807

	return sum8, sum16, sum32, sum64
}
