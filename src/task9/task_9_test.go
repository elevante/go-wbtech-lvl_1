package main

import (
	"go-wbtech-lvl_1/pkg"
	"testing"
)

func TestConveyor(t *testing.T) {
	testTable := []struct {
		numbers []int
		expect  []int
	}{
		{
			numbers: []int{2, 4, 8},
			expect:  []int{4, 8, 16},
		},
		{

			numbers: []int{2000, 4000, 8000},
			expect:  []int{4000, 8000, 16000},
		},
		{
			numbers: []int{-2, -4, -8},
			expect:  []int{-4, -8, -16},
		},
		{
			numbers: []int{},
			expect:  []int{},
		},
	}

	for _, testCase := range testTable {
		multiplyByTwo := conveyor(testCase.numbers)

		if pkg.CompareArrays(multiplyByTwo, testCase.expect) == false {
			t.Errorf("Incorrect result. Expect %d, got%d",
				testCase.expect,
				multiplyByTwo)

		}

	}
}
