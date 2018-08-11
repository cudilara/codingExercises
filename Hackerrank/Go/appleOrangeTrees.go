package main

import (
	"fmt"
)

func getValues() ([]int, []int, []int, []int) {
	house := make([]int, 2)
	for i := 0; i < 2; i++ {
		fmt.Scan(&house[i])
	}

	trees := make([]int, 2)
	for i := 0; i < 2; i++ {
		fmt.Scan(&trees[i])
	}

	fruitNum := make([]int, 2)
	for i := 0; i < 2; i++ {
		fmt.Scan(&fruitNum[i])
	}

	appleDist := make([]int, fruitNum[0])
	for i := 0; i < fruitNum[0]; i++ {
		fmt.Scan(&appleDist[i])
	}

	orangeDist := make([]int, fruitNum[1])
	for i := 0; i < fruitNum[1]; i++ {
		fmt.Scan(&orangeDist[i])
	}
	return house, trees, appleDist, orangeDist
}

func getFruitPosition(trees []int, appleDist []int, orangeDist []int) ([]int, []int){
	var applePos []int
	var orangePos []int
	var pos int
	for apple := range appleDist {
		pos = trees[0] + appleDist[apple]
		applePos = append(applePos, pos)
	}
	for orange := range orangeDist {
		pos = trees[1] + orangeDist[orange]
		orangePos = append(orangePos, pos)
	}
	return applePos, orangePos
}

func countHouseFruit(house []int, applePos []int, orangePos []int) (int, int){
	var appleN int
	var orangeN int
	for i := range applePos {
		if applePos[i] >= house[0] {
			if applePos[i] <= house[1] {
				appleN += 1
			}
		}
	}
	for k := range orangePos {
		if orangePos[k] >= house[0] {
			if orangePos[k] <= house[1]{
				orangeN += 1
			}
		}
	}
	return appleN, orangeN
}

func main() {
	house, trees, appleDist, orangeDist := getValues()
	applePos, orangePos := getFruitPosition(trees, appleDist, orangeDist)
	appleN, orangeN := countHouseFruit(house, applePos, orangePos)
	fmt.Println(appleN)
	fmt.Println(orangeN)
}