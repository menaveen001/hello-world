package main

import "fmt"

func main() {
	fmt.Println("hello world")
	addition()
	substraction()
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
