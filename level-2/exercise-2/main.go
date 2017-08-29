package main

import "fmt"

func main() {
	a := (42 == 42)
	b := (42 <= 42)
	c := (43 >= 42)
	d := (43 != 42)
	e := (41 < 42)
	f := (43 > 42)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
