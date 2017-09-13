package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	p1 := person{
		first: "Brent",
		last:  "Biglin",
		iceCream: []string{
			"vanilla",
			"pistachio",
		},
	}
	p2 := person{
		first: "Whitney",
		last:  "Chappell",
		iceCream: []string{
			"chocolate",
			"fudge",
		},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	for i, v := range p1.iceCream {
		fmt.Println(i, v)
	}
	for i, v := range p2.iceCream {
		fmt.Println(i, v)
	}
}
