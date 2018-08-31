package main

func rotateLeft(arr []int32, d int32) []int32 {
	arrSize := len(arr)
	shiftValue := int(d)
	var res = make([]int32, arrSize)
	var currentVal int32
	shiftValue = normalizeShiftValue(arrSize, shiftValue)
	if arrSize == shiftValue {
		return arr
	}
	for i := 0; i < arrSize; i++ {
		currentVal = arr[i]
		if i < shiftValue {
			res[arrSize-shiftValue+i] = currentVal
		} else {
			res[i-shiftValue] = currentVal
		}
	}
	return res
}

func normalizeShiftValue(arrSize, shiftValue int) int {
	if arrSize < shiftValue {
		shiftValue = shiftValue % arrSize
	}
	return shiftValue
}
