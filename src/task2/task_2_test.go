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

func TestOutputSquares(t *testing.T) {
	testTable := []struct {
		numbers []int
		expect  []int
	}{
		{
			numbers: []int{4, 6, 8, 9, 10, 20},
			expect:  []int{16, 36, 64, 81, 100, 400},
		},
		{

			numbers: []int{500, 600, 800, 900, 100, 200},
			expect:  []int{250000, 360000, 640000, 810000, 10000, 40000},
		},
		{
			numbers: []int{-4, -6, 8, 9, -10, 20},
			expect:  []int{16, 36, 64, 81, 100, 400},
		},
		{
			numbers: []int{},
			expect:  []int{},
		},
	}

	for _, testCase := range testTable {
		squares := outputSquares(testCase.numbers)

		if pkg.CompareArrays(squares, testCase.expect) == false {
			t.Errorf("Incorrect result. Expect %d, got%d",
				testCase.expect,
				squares)

		}
	}
}
