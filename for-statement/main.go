package main

import "fmt"

func main() {
	x := 1
	for x < 11 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")

	for y := 1; y < 11; y++ {
		fmt.Println(y)
	}
	fmt.Println("done.")

	z := 0
	for {
		z++
		if z <= 10 {
			fmt.Println(z)
		}
		if z > 10 {
			break
		}
	}
	fmt.Println("done.")
}
