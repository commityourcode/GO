package main

import "fmt"

//create own type
type hotdog int

var y int

var x hotdog

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	//conversion
	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
