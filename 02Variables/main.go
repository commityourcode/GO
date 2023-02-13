package main

import "fmt"

var JwtToken string = "xxxxx"

const PI float64 = 3.14

func main() {
	var username string = "Anurag"
	fmt.Println(username)
	fmt.Printf("Variable is of Type: %T \n", username)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("Variable is of Type: %T \n", isLogged)

	var val int = 256
	fmt.Println(val)
	fmt.Printf("Variable is of Type: %T \n", val)

	var numver float64 = 255.345678123
	fmt.Println(numver)
	fmt.Printf("Variable is of Type: %T \n", numver)

	//Default value & Alias
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of Type: %T \n", anothervariable)

	//No var style walrus operator
	numberofUser := 3000
	fmt.Println(numberofUser)
	fmt.Printf("Variable is of Type: %T \n", numberofUser)

	fmt.Println(JwtToken)
	fmt.Println(PI)
}
