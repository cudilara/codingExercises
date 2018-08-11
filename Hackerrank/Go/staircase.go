package main

import (
	"fmt"
	"strings"
)

func main() {
	space := " "
	symbol := "#"
	var printSpace = ""
	var printSymbol = ""
	var N int
	fmt.Scanln(&N)
	for i := 0; i <= N; i++ {
		printSpace = strings.Repeat(space, N - i)
		fmt.Print(printSpace)
		printSymbol = strings.Repeat(symbol, i)
		fmt.Println(printSymbol)
	}
}

