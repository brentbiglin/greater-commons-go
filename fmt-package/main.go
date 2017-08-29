package main

import "fmt"

var a int
var b string = "James Bond"

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	s := fmt.Sprint(a, " something more ", b)
	fmt.Println(s)
	s2 := fmt.Sprintf("%v\t%T\t%T\n", "to pass in", a, b)
	fmt.Println(s2)
}
