package main

import "fmt"

func main() {
	fmt.Println("hello world")
	addition()

}
func addition() {

	a, b := 2, 3
	sum := a + b
	fmt.Println(sum)
}
