package main

import "fmt"

func compArr(a []int, b []int) (int, int){
	Ascore, Bscore := 0, 0
	for i := 0; i < 3; i++ {
		if a[i] > b[i]{
			Ascore += 1
		} else if a[i] < b[i] {
			Bscore += 1
		}
	}
	return Ascore, Bscore
}

func main() {
	a := make([]int, 3)
	b := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&a[i])
	}
	for j := 0; j < 3; j++ {
		fmt.Scan(&b[j])
	}
	resA, resB := compArr(a, b)
	fmt.Println(resA, resB)
}