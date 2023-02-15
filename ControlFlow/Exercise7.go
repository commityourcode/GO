package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "James Bond" {
		fmt.Println(x)
	} else if x == "Money Penny" {
		fmt.Println("Money Penny")
	} else {
		fmt.Println("None")
	}
}
