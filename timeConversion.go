package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	format := s[len(s)-2:]
	timeArr := strings.Split(s[:len(s)-2], ":")
	h, m, sc := timeArr[0], timeArr[1], timeArr[2]
	if format == "PM" {
		if h != "12" {
			hour, _ := strconv.Atoi(h)
			hour += 12
			h = strconv.Itoa(hour)
		}
	} else {
		if h == "12" {
			h = "00"
		}
	}

	return fmt.Sprintf("%s:%s:%s", h, m, sc)
}

func main() {
	s := "07:05:45PM"
	fmt.Println(timeConversion(s))
}
