package acdc

//Sum adds an unlimited number of values of type int
func Sum(xi ...int) int {
	sum := 0
	for _, y := range xi {
		sum += y
	}
	return sum
}
