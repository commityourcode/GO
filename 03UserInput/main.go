package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating for Pizza:")
	input, _ := reader.ReadString('\n')

	// comma ok // err err

	fmt.Println("Thanks for Rating", input)
	fmt.Printf("Type of this Rating : %T \n", input)
}
