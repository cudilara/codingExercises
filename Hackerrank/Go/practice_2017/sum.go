package practice_2017

import "fmt"

func addArr(a []int) int {
	sum := 0
	for i := range a {
		sum += a[i]
	}
	return sum
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var N int
	fmt.Scanln(&N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	res := addArr(a)
	fmt.Println(res)
}
