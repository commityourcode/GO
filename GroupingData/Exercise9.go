package main

import "fmt"

//create a map with key type string and value type string and add a record
func main() {

	m := map[string]int{"Jan": 1, "Feb": 2, "Mar": 3}
	fmt.Println(m)
	m["April"] = 4
	for k, v := range m {
		fmt.Println(k, v)
	}
}
