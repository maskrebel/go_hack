package main

import (
	"fmt"
	"strings"
)

func pangrams(s string) string {
	s = strings.ToLower(s)
	mapChar := make(map[string]struct{})
	for _, v := range s {
		mapChar[string(v)] = struct{}{}
	}
	if len(mapChar) == 27 {
		return "pangram"
	}
	return "not pangram"
}

func main() {
	s := "We promptly judged antique ivory buckles for the next prize"
	fmt.Println(pangrams(s))
}
