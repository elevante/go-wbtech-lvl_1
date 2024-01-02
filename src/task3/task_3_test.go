package main

import (
	"testing"
)

func TestSquare(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		t.Parallel()
		t.Log("simple")
		var x = 2
		expect := 4

		result := square(x)

		if result != expect {
			t.Errorf("Incorrect result. Expect %d, got%d",
				result,
				expect)
		}
	})
	t.Run("medium", func(t *testing.T) {
		t.Parallel()
		t.Log("medium")
		var x = 550
		expect := 302500

		result := square(x)

		if result != expect {
			t.Errorf("Incorrect result. Expect %d, got%d",
				result,
				expect)
		}
	})
	t.Run("negative", func(t *testing.T) {
		t.Parallel()
		t.Log("negative")
		var x = -5
		expect := 25

		result := square(x)

		if result != expect {
			t.Errorf("Incorrect result. Expect %d, got%d",
				result,
				expect)
		}
	})
}

func TestOutputSumSquares(t *testing.T) {
	testTable := []struct {
		numbers []int
		expect  int
	}{
		{
			numbers: []int{4, 6, 8, 9, 10, 20},
			expect:  697,
		},
		{

			numbers: []int{500, 600, 800, 900, 100, 200},
			expect:  2110000,
		},
		{
			numbers: []int{-4, -6, 8, 9, -10, 20},
			expect:  697,
		},
		{
			numbers: []int{},
			expect:  0,
		},
	}

	for _, testCase := range testTable {
		sumSquares := outputSumSquares(testCase.numbers)

		if sumSquares != testCase.expect {
			t.Errorf("Incorrect result. Expect %d, got%d",
				testCase.expect,
				sumSquares)

		}
	}
}
