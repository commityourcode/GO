package main

import "testing"

// benchmark the speed of computing the 10th number in the Fibonacci series.
func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
