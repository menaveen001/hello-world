package main

import "fmt"

func main() {
	fmt.Println("hello world")
	addition()
	substraction()
	multipiction()
}
func addition() {

	a, b := 2, 3
	sum := a + b
	fmt.Println(sum)
}
func substraction() {

	a, b := 8, 5
	sub := a - b
	fmt.Println(sub)
}

func multipiction() {

	a, b := 8, 5
	multi := a * b
	fmt.Println(multi)
}
