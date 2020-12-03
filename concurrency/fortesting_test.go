package main

import "testing"

func BenchmarkF2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		f2()
	}
}

func BenchmarkF1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		f1()
	}
}

// 0.00299
