package main

import "fmt"

func main() {
	fmt.Println("hello world")
	addition()
	substraction()
	multipiction()
	getUserInput()
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

	a, b := 10, 5
	multi := a * b
	fmt.Println(multi)
}
func getUserInput() {
	var FirstName string
	var LastName string
	var age int

	fmt.Println("Enter your first name")
	fmt.Scan(&FirstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&LastName)

	fmt.Println("Enter your age")
	fmt.Scan(&age)

}
