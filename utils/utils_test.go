package utils

import (
	"github.com/SheltonZhu/tools/functools"
	"testing"
)

func TestTypeCheck(t *testing.T) {
	TypeCheck(13, -14.3, "BELGIUM", complex(1, 2), nil, false)
}

func TestRangeIntsSlice(t *testing.T) {
	expected := []int{0, 2, 4, 6, 8}
	start, end, step := 0, 10, 2
	actual := RangeIntsSlice(start, end, step)
	if !functools.CompareSlice(functools.IntSlice(actual), functools.IntSlice(expected)) {
		t.Errorf("RangeIntsSlice(%v, %v, %v) = %v; expected %v", start, end, step, actual, expected)
	}
}

type TestSet struct {
	In       int
	Expected int
}

func TestDaysFromYear(t *testing.T) {
	testSets := []TestSet{
		{2008, 366},
		{2020, 366},
		{2021, 365},
	}

	for _, testSet := range testSets {
		actual := DaysFromYear(testSet.In)
		if actual != testSet.Expected {
			t.Errorf("DaysFromYear(%v) = %v; expected %v", testSet.In, actual, testSet.Expected)
		}
	}
}
