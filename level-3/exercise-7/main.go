package main

import "fmt"

var z string = "Yoda"

func main() {
	x := "I'm a Jedi, like my father before me."
	y := "If you only knew the power of the dark side."
	if z == "Luke Skywalker" {
		fmt.Println(x)
	} else if z == "Darth Vader" {
		fmt.Println(y)
	} else {
		fmt.Println("Do or do not, there is no try.")
	}
}
