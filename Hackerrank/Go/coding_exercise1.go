package main

import "strings"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string, K int) string {
	// write your code in Go 1.4
	newStr := strings.Replace(S, "-", "", -1)
	count := K
	retString := ""
	runes := []rune(newStr)
	for i := len(runes) - 1; i >= 0; i-- {
		if count == 0 {
			retString = "-" + retString
		}
		retString = string(runes[i]) + retString
		count = count - 1
	}
	return retString
}
