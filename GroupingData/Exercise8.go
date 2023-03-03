package main

import "fmt"

//create a mapwith key type string and value type string
func main() {

	m := map[string]int{"Jan": 1, "Feb": 2, "Mar": 3}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
