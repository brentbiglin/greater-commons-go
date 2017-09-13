package main

import "fmt"

func main() {
	b := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Brent",
		friends: map[string]int{
			"Jason":   555,
			"Eddie":   666,
			"Whitney": 777,
		},
		favDrinks: []string{
			"Coffee",
			"Beer",
		},
	}
	fmt.Println(b.first)
	fmt.Println(b.friends)
	fmt.Println(b.favDrinks)

	for k, v := range b.friends {
		fmt.Println(k, v)
	}

	for i, val := range b.favDrinks {
		fmt.Println(i, val)
	}
}
