package main

import "fmt"

func main() {
	//Declaration for slice same as array but we don't have to define size here
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//Slicing a slice
	fmt.Println(x[0:5])  //42-46
	fmt.Println(x[5:10]) //47-51
	fmt.Println(x[2:7])  //44-48
	fmt.Println(x[1:6])  //43-47

}
