package practice_2017

import "fmt"

func main() {
	result := "NO"
	var value int
	a := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Scan(&a[i])
	}
	locdiff := a[0] - a[2]
	speeddiff := a[1] - a[3]
	if speeddiff != 0 {
		value = locdiff % speeddiff
	} else {
		value = 1
	}
	if locdiff == 0 && speeddiff == 0 {
		result = "YES"
	}
	if (locdiff < 0 && speeddiff > 0) || (locdiff > 0 && speeddiff < 0) {
		if value == 0 {
			result = "YES"
		}
	}
	fmt.Println(result)
}
