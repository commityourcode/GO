package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizaa app")
	fmt.Println("Please rate our pizza between 1 and 5")

	read := bufio.NewReader(os.Stdin)

	input, _ := read.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to rating: ", numRating+1)
	}
}
