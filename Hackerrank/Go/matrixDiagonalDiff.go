package main

import (
	"fmt"
	"math"
)

func sumDiagonals(a [][]int, N int) (int, int){
	var sum1 int = 0
	var sum2 int = 0
	for i := 0; i < N; i++ {
		sum1 += a[i][i]
		sum2 += a[i][len(a) - 1 - i]
	}
	return sum1, sum2
}

func getDifference(a float64, b float64) float64{
	return math.Abs(a - b)
}

func main() {
	var N int
	fmt.Scanln(&N)
	a := make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, N)
		for k:= 0; k < N; k++{
			fmt.Scan(&a[i][k])
		}
	}
	sum1, sum2 := sumDiagonals(a, N)
	res := getDifference(float64(sum1), float64(sum2))
	fmt.Println(res)
}
