package main

import "fmt"

func main() {
	favSport := "basketball"
	switch favSport {
	case "baseball":
		fmt.Println("Go Indians!")
	case "football":
		fmt.Println("Go Bucks!")
	case "hockey":
		fmt.Println("Go Bluejackets!")
	case "basketball":
		fmt.Println("Go Spurs Go!")
	}
}
