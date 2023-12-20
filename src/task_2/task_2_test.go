package main

import (
	"go-wbtech-lvl_1/src/pkg"
	"testing"
)

func TestSquare(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		t.Parallel()
		t.Log("simple")
		var x = 2
		expected := 4

		result := square(x)

		if result != expected {
			t.Errorf("Incorrect result. Expect %d, got%d",
				result,
				expected)
		}
	})
	t.Run("medium", func(t *testing.T) {
		t.Parallel()
		t.Log("medium")
		var x = 550
		expected := 302500

		result := square(x)

		if result != expected {
			t.Errorf("Incorrect result. Expect %d, got%d",
				result,
				expected)
		}
	})
	t.Run("negative", func(t *testing.T) {
		t.Parallel()
		t.Log("negative")
		var x = -5
		expected := 25

		result := square(x)

		if result != expected {
			t.Errorf("Incorrect result. Expect %d, got%d",
				result,
				expected)
		}
	})
}

func TestOutputSquares(t *testing.T) {
	testTable := []struct {
		numbers  []int
		expected []int
	}{
		{
			numbers:  []int{4, 6, 8, 9, 10, 20},
			expected: []int{16, 36, 64, 81, 100, 400},
		},
		{

			numbers:  []int{500, 600, 800, 900, 100, 200},
			expected: []int{250000, 360000, 640000, 810000, 10000, 40000},
		},
		{
			numbers:  []int{-4, -6, 8, 9, -10, 20},
			expected: []int{16, 36, 64, 81, 100, 400},
		},
		{
			numbers:  []int{},
			expected: []int{},
		},
	}

	for _, testCase := range testTable {
		squares := outputSquares(testCase.numbers)

		if pkg.CompareArrays(squares, testCase.expected) == false {
			t.Errorf("Incorrect result. Expected %d, got%d",
				testCase.expected,
				squares)

		}
	}
}
