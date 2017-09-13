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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
	}
}
