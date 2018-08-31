package main

import (
	"reflect"
	"testing"
)

type testArray struct {
	array      []int32
	shiftValue int32
	expected   []int32
}

var testArrays = []testArray{
	{array: []int32{23, 45, 67}, shiftValue: 1, expected: []int32{45, 67, 23}},
	{array: []int32{0}, shiftValue: 1, expected: []int32{0}},
	{array: []int32{23, 45, 67}, shiftValue: 3, expected: []int32{23, 45, 67}},
	{array: []int32{23, 45, 67, 12, 98, 0}, shiftValue: 27, expected: []int32{12, 98, 0, 23, 45, 67}},
}

func TestRotateLeft(t *testing.T) {
	for _, testArr := range testArrays {
		actual := rotateLeft(testArr.array, testArr.shiftValue)
		if !reflect.DeepEqual(testArr.expected, actual) {
			t.Errorf("left shift array %v: expected %d, actual %d", testArr.array, testArr.expected, actual)
		}
	}
}
