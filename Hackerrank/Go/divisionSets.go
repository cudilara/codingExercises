package main

import (
	"fmt"
)

func inputValues() ([]int, []int) {
	nm := make([]int, 2)
	for i := 0; i < 2; i++ {
		fmt.Scan(&nm[i])
	}

	aset := make([]int, nm[0])
	for i := 0; i < nm[0]; i++ {
		fmt.Scan(&aset[i])
	}

	bset := make([]int, nm[1])
	for i := 0; i < nm[1]; i++ {
		fmt.Scan(&bset[i])
	}
	return aset, bset
}

func checkA(n int, set []int) bool {
	for i := range set {
		if n % set[i] != 0 {
			return false
		}
	}
	return true
}

func checkB(n int, set []int) bool {
	for i := range set {
		if n != 0 {
			if set[i] % n != 0 {
				return false
			}
		}
	}
	return true
}

func checkLastAValue(aset []int, bset []int) bool {
	return checkA(aset[len(aset) - 1], aset) && checkB(aset[len(aset) - 1], bset)
}

func getValsBetweenSets(aset []int, bset []int) int {
	rNum := 0
	for n := aset[len(aset) - 1] + 1; n < bset[0]; n++ {
		if checkA(n, aset) && checkB(n, bset) {
			rNum += 1
		}
	}
	return rNum
}

func main() {
	resNum := 0
	statusA := false
	statusB := false
	aset, bset := inputValues()
	if checkLastAValue(aset, bset) {
		resNum += 1
	}
	resNum += getValsBetweenSets(aset, bset)
	for i := range bset {
		statusA = checkA(bset[i], aset)
		statusB = checkB(bset[i], bset)
		if statusA && statusB {
			resNum += 1
		}
	}
	fmt.Println(resNum)
}
