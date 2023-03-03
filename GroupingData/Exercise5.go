package main

import "fmt"

//To delete from slices we use append along with slicing
//42,43,44,48,49,50,51
func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	//x = append(x, x...) //append each value of x
	x = append(x[:3], x[6:]...) // value less than 4th index and 6th to upto last index
	// x = append(x[0:3], x[6]...)
	fmt.Println(x)

}
