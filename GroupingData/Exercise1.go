package main

import "fmt"

func main() {
	//declare array of 5 no
	x := [5]int{42, 43, 44, 45, 46}
	for i := 0; i < 5; i++ {
		fmt.Println(x[i])
	}
	//indices and values
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
