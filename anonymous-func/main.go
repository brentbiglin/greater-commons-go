package main

import "fmt"

func main() {
	foo()

	// This is an anonymous function.
	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)
}

func foo() {
	fmt.Println("foo ran")
}
