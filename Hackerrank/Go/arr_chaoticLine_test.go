package main

import (
	"reflect"
	"testing"
)

type testValArrays struct {
	givenArr []int32
	expected int
}

var tests = []testValArrays{
	{givenArr: []int32{2}, expected: 0},
	{givenArr: []int32{1, 2, 3}, expected: 0},
	{givenArr: []int32{2, 1, 3}, expected: 1},
	{givenArr: []int32{2, 1, 5, 3, 4}, expected: 3},
	{givenArr: []int32{2, 5, 1, 3, 4}, expected: -1},
	{givenArr: []int32{5, 1, 2, 3, 7, 8, 6, 4}, expected: -1},
	{givenArr: []int32{1, 2, 5, 3, 7, 8, 6, 4}, expected: 7},
	{givenArr: []int32{1, 2, 5, 3, 4, 7, 8, 6}, expected: 4},
}

func TestFindMinimum(t *testing.T) {
	for _, arr := range tests {
		actual, _ := minimumSwitchesInitial(arr.givenArr)
		if !reflect.DeepEqual(arr.expected, actual) {
			t.Errorf("error in array %v: expected %d, actual %d", arr.givenArr, arr.expected, actual)
		}
	}
}
