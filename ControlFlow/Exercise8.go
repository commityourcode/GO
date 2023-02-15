package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Should not be able to")
	case true:
		fmt.Println("able to connect")
	}
}
