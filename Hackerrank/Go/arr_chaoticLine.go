package main

import "math"

// with solution from Hackerrank discussion
func minimumSwitches(q []int32) (int, string) {
	var switches, currentVal, switchNum int
	switched := false
	var temp int32
	for i := 0; i < (len(q) - 1); i++ {
		currentVal = int(q[i])
		if int(math.Abs(float64(currentVal))) > (i + 3) {
			return -1, "Too chaotic"
		}
		for j := 0; j < (len(q) - 1); j++ {
			if q[j] > q[j+1] {
				temp = q[j]
				q[j] = q[j+1]
				q[j+1] = temp
				switches += 1
				switched = true
			}
		}
		if switched {
			switched = false
		} else {
			break
		}
	}
	return switchNum, ""
}

func minimumSwitchesInitial(q []int32) (int, string) {
	var currentVal, switchNum int
	minSwitches := 0
	if len(q) < 2 {
		return minSwitches, ""
	}
	for i := 0; i < len(q); i++ {
		currentVal = int(q[i])
		if int(math.Abs(float64(currentVal))) > (i + 3) {
			return -1, "Too chaotic"
		}
		if currentVal != (i + 1) {
			switchNum += 1
		}
	}
	if switchNum%2 != 0 {
		switchNum = switchNum/2 + 1
	} else {
		switchNum = switchNum / 2
	}

	return switchNum, ""
}
