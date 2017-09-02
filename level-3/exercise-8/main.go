package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("nope")
	case true:
		fmt.Println("yep")
	}
}
