package main

func arrToMap(ar []int32) map[int32]int32 {
	myMap := make(map[int32]int32)
	for i := 0; i < len(ar); i += 1 {
		if val, ok := myMap[ar[i]]; ok {
			myMap[ar[i]] = val + 1
		} else {
			myMap[ar[i]] = 1
		}
	}
	return myMap
}

func getMatchedPairs(n int32, ar []int32) int32 {
	var pair = 0
	var res = 0
	myMap := arrToMap(ar)
	for _, value := range myMap {
		pair = int(value / 2)
		res += pair
	}
	return int32(res)
}
