package main

import (
	"fmt"
	"sort"
)

func getMaxCount(a []int) int{
	cnt := 0
	max := a[len(a) - 1]
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == max {
			cnt += 1
		}
	}
	return cnt
}

func main() {
	var N int
	fmt.Scanln(&N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	blown := getMaxCount(a)
	fmt.Println(blown)
}

