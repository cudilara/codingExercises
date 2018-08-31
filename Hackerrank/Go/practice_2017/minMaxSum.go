package practice_2017

import (
	"fmt"
	"sort"
)

func findArrays(nums []int) ([]int, []int) {
	var small []int
	var large []int
	sort.Ints(nums)
	small = nums[:len(nums)-1]
	large = nums[1:]
	return small, large
}

func getSums(smallNums []int, largeNums []int) (int64, int64) {
	var min, max int64
	for i := range smallNums {
		min += int64(smallNums[i])
	}
	for i := range largeNums {
		max += int64(largeNums[i])
	}
	return min, max
}

func main() {
	nums := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&nums[i])
	}
	smallNums, largeNums := findArrays(nums)
	small, large := getSums(smallNums, largeNums)
	fmt.Println(small, large)
}
