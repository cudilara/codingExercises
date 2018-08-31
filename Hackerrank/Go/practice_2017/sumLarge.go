package practice_2017

import "fmt"

func addLargeArr(a []int64) int64 {
	var sum int64 = 0
	for i := range a {
		sum += a[i]
	}
	return sum
}

func main() {
	var N int
	fmt.Scanln(&N)
	a := make([]int64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	res := addLargeArr(a)
	fmt.Println(res)
}
