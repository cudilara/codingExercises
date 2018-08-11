package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var usTime string
	var mhr int
	fmt.Scanln(&usTime)
	parsed := strings.Split(usTime, ":")
	ushr := parsed[0]
	min := parsed[1]
	secAmPm := parsed[2]
	sec := secAmPm[:len(secAmPm)-2]
	amPm := secAmPm[2:]
	if amPm == "PM" {
		if ushr == "12" && min == "00" && sec == "00" {
			ushr = "12"
		} else {
			num, _ := strconv.Atoi(ushr)
			mhr = num + 12
			if mhr > 23 {
				ushr = "00"
			}
			ushr = strconv.Itoa(mhr)
		}
	}
	if amPm == "AM" && ushr == "12" && min == "00" && sec == "00"{
		ushr = "00"
	}
	result := ushr + ":" + parsed[1] + ":" + sec
	fmt.Println(result)
}