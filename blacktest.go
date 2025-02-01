// package main

import (
	"testing"
)

func Add(a,b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	tests := []struct {
		input1, input2, expected int
}{
	{2, 3, 5},
	{0, 0, 0},
	{-1, 1, 0},
}

for _, tt := range tests {
	result := Add(tt.input1, tt.input2)
	if result != tt.expected {
		t.Errorf("Add(%d, %d) = %d; want %d", tt.input1, tt.input2, result, tt.expected)
	}
}
}