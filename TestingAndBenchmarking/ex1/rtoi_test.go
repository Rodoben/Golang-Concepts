package main

import "testing"
import "fmt"

func TestRtoi(t *testing.T) {
	X := romanToInt("IV") //4
	fmt.Println(X)
	if X != 4 {
		t.Error("Expected 4 got", X)
	}
}
