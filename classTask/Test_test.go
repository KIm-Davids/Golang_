package main

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Add(firstNumber int, secondNumber int) int {
	result := firstNumber + secondNumber
	return result
}

func TestMyCode(t *testing.T) {
	result := Add(5, 5)
	expected := 15

	//if result != expected {
	//	t.Errorf("Add(5, 5) = %d; want %d", result, expected)
	//}

	assert.(t, result, expected)
}
