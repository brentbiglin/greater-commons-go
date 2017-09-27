package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // & gives you the address

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an address

	fmt.Println(*&a)
}
