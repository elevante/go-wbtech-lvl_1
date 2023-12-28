package main

import (
	"go-wbtech-lvl_1/pkg"
	"testing"
)

func TestIntersectionOfManya(t *testing.T) {
	testTable := []struct {
		numbersFirst  []int
		numbersSecond []int
		expect        []int
	}{
		{
			numbersFirst:  []int{2, 4, 8, 10},
			numbersSecond: []int{1, 3, 4, 10},
			expect:        []int{4, 10},
		},
		{
			numbersFirst:  []int{2, -4, -8, -10},
			numbersSecond: []int{1, 3, -4, -10},
			expect:        []int{-4, -10},
		},
		{
			numbersFirst:  []int{},
			numbersSecond: []int{},
			expect:        []int{},
		},
	}

	for _, testCase := range testTable {
		result := intersectionOfManya(testCase.numbersFirst, testCase.numbersSecond)

		if pkg.CompareArrays(result, testCase.expect) == false {
			t.Errorf("Incorrect result. Expect %d, got%d",
				testCase.expect,
				result)
		}

	}
}
