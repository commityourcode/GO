package main

func main() {
	// fmt.Println("2 + 3 =", mySum(2, 3))
	// fmt.Println("4 + 7 =", mySum(4, 7))
	// fmt.Println("5 + 9 =", mySum(5, 9))
}

func mySum(xi ...int) int {
	sum := 0
	for _, y := range xi {
		sum += y
	}
	return sum
}
