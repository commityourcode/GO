package main

import "fmt"

//create a slice of slice of strings([][]string) store in multi-dimensional slice
func main() {

	x := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}
	y := []string{"01", "02", "03", "04", "05", "06"}
	fmt.Println(x)
	fmt.Println(y)

	//slice of slice of string
	xy := [][]string{x, y}
	fmt.Println(xy)

	//Range it over slice of slice of string
	for i, xy1 := range xy {
		fmt.Println("Record:", i)
		for j, val := range xy1 {
			fmt.Printf("\t Index position: %v \t value: \t %v\n", j, val)
		}
	}

}
