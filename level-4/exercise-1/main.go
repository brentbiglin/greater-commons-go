package main

import "fmt"

func main() {
	x := [5]int{42, 44, 45, 43, 41}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
