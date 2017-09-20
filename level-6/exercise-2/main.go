package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...)
	fmt.Println(n)

	// no need to unfurl a slice
	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2 := bar(ii2)
	fmt.Println(n2)
}

// this is a variadic parameter
func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// this is a slice of int, so no need to unfurl
func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
