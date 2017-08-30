package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("this should print")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, should print")
	}

	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	default:
		fmt.Println("this is the default")
	}
}
