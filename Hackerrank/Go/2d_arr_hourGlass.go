package main

import "fmt"

func hourglassSum(arr [][]int32) (int32, error) {
	var largest, current, topVal, midVal, bottomVal int32
	ret, err := checkArraySize(arr)
	if err != nil {
		return ret, err
	}
	largest = initializeLargestValue(arr)
	for i := 0; i < (len(arr) - 2); i++ {
		for j := 0; j < (len(arr[i]) - 2); j++ {
			topVal = arr[i][j] + arr[i][j+1] + arr[i][j+2]
			midVal = arr[i+1][j+1]
			bottomVal = arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			current = topVal + midVal + bottomVal
			if current > largest {
				largest = current
			}
		}
	}
	return largest, nil
}

func checkArraySize(arr [][]int32) (int32, error) {
	if len(arr) < 3 {
		return -1, fmt.Errorf("array has less than 3 rows")
	}
	return 0, nil
}

func initializeLargestValue(arr [][]int32) int32 {
	return arr[0][0] + arr[0][1] + arr[0][2] + arr[1][1] + arr[2][0] + arr[2][1] + arr[2][2]
}
