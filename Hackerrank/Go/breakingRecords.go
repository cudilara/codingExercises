package main

import "fmt"

func main() {
	var N int
	var resMin = 0
	var resMax = 0
	fmt.Scanln(&N)
	scores := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&scores[i])
	}
	var minR = scores[0]
	var maxR = scores[0]
	for g := range scores {
		if scores[g] > maxR {
			maxR = scores[g]
			resMax += 1
		}
		if scores[g] < minR {
			minR = scores[g]
			resMin += 1
		}
	}
	fmt.Println(resMax, resMin)
}
