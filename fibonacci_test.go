package main

import "testing"

func BenchmarkFib39(b *testing.B) {
	fib(39)
}

func BenchmarkFibTCO_8388608(b *testing.B) {
	fibTCO(8388608, 0, 1)
}
func BenchmarkFibIter(b *testing.B) {
	fibIter(536870912)
}
