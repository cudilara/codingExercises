package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution2(A []int) int {
	fruitMap := make(map[int]int)
	value := 0
	for i := 0; i < len(A); i++ {
		if _, ok := fruitMap[A[i]]; ok {
			value = fruitMap[A[i]] + 1
			fruitMap[A[i]] = value
		} else {
			fruitMap[A[i]] = A[i]
		}
	}
	print(fruitMap)
	return 2
}
