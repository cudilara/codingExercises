package main

import (
	"testing"
)

type testArr struct {
	arr      [][]int32
	expected int32
}

var testCases = []testArr{
	{arr: [][]int32{{1, 1}}, expected: -1},
	{arr: [][]int32{{1, 1, 1},
		{0, 0}}, expected: -1},
	{arr: [][]int32{nil}, expected: -1},
	{arr: [][]int32{{1, 1, 1, 0},
		{0, 0, 0, 0},
		{1, 1, 1, 0}}, expected: 6},
	{arr: [][]int32{{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0}}, expected: 7},
	{arr: [][]int32{{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, -9, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 1, 2, 4, 0}}, expected: 28},
}

func TestCalculateSums(t *testing.T) {
	for _, testArr := range testCases {
		actual, _ := hourglassSum(testArr.arr)
		if actual != testArr.expected {
			t.Errorf("hour glass %v: expected %d, actual %d", testArr.arr, testArr.expected, actual)
		}
	}
}
