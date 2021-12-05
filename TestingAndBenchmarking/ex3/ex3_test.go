package main

import "testing"

func TestTableCalculate(t *testing.T) {
	type test struct {
		input    []int
		expected int
	}
	tests := []test{
		test{[]int{2, 6}, 8},
		test{[]int{1, 2}, 2},
		test{[]int{2, 6}, 8},
		test{[]int{2, 6}, 8},
		test{[]int{2, 6}, 8},
	}

	for _, test := range tests {
		if output := Calculate(test.input...); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func BenchmarkCalculate(b *testing.B) {
	type test struct {
		input    []int
		expected int
	}
	tests := []test{
		test{[]int{2, 6}, 8},
		test{[]int{1, 1}, 3},
		test{[]int{2, 6}, 8},
		test{[]int{2, 6}, 8},
		test{[]int{2, 6}, 8},
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Calculate(test.input...)
		}
	}
}
