package main

import "fmt"

func main() {
	bd := 1996
	for {
		if bd <= 2023 {
			fmt.Println(bd)
		}
		bd++
	}
}
