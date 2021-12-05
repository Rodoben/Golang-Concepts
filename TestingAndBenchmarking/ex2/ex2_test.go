package main

import "testing"

func TestEx2(t *testing.T) {
	x := mySum(2, 3, 1)
	if x != 6 {
		t.Error("Expected 5 but got", x)
	}
}

func BenchmarkTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mySum(2, 3, 1)
	}
}
