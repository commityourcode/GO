package main

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome My Dear James" {
		t.Error("got", s, "Expected", "Welcome My Dear James")
	}
}

func ExampleGreet() {
	fmt.Println("")
	fmt.Println(Greet("James"))
	//Output : Welcome My Dear James
}

func BenchmarkGreet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Greet("James")
	}
}
