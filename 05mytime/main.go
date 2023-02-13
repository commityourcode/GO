package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time in G Lang")
	presentTime := time.Now()

	fmt.Println("Current Time: ", presentTime)
	fmt.Println("Current Time: ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.February, 8, 16, 3, 0, 0, time.Now().UTC().Location())

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
