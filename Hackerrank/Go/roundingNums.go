package main

import (
	"fmt"
)

func calcGrade(a []int) []int {
	var result []int
	var leftover int
	var divider int
	var grade int
	var nextMultiple int
	for i := range a {
		if a[i] > 37{
			divider = a[i] / 5
			nextMultiple = (divider + 1) * 5
			leftover = nextMultiple - a[i]
			if leftover < 3 {
				grade = nextMultiple
			} else {
				grade = a[i]
			}
		} else {
			grade = a[i]
		}
		result = append(result, grade)
	}
	return result
}

func main() {
	var N int
	fmt.Scanln(&N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	result := calcGrade(a)
	for i := range result {
		fmt.Println(result[i])
	}
}