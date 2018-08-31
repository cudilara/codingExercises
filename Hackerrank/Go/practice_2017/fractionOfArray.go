package practice_2017

import "fmt"

func getFractions(a []int, N int) (float64, float64, float64) {
	positive := 0
	negative := 0
	zero := 0
	for i := range a {
		if a[i] < 0 {
			negative += 1
		} else if a[i] > 0 {
			positive += 1
		} else {
			zero += 1
		}

	}
	posFrac := float64(positive) / float64(N)
	negFrac := float64(negative) / float64(N)
	zerFrac := float64(zero) / float64(N)
	return posFrac, negFrac, zerFrac
}

func main() {
	var N int
	fmt.Scanln(&N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	d, b, c := getFractions(a, N)
	fmt.Printf("%.6f\n%.6f\n%.6f", d, b, c)
}
