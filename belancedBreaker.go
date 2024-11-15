package main

import (
	"bufio"
	"fmt"
	"os"
)

func isBalanced(s string) string {
	mapBalance := map[string]string{"}": "{", "]": "[", ")": "("}
	tmpBracket := make([]string, 0)
	for _, v := range s {
		if spouse, ok := mapBalance[string(v)]; ok {
			if len(tmpBracket) == 0 || tmpBracket[len(tmpBracket)-1] != spouse {
				return "NO"
			} else {
				tmpBracket = tmpBracket[:len(tmpBracket)-1]
			}
		} else {
			tmpBracket = append(tmpBracket, string(v))
		}
	}
	if len(tmpBracket) > 0 {
		return "NO"
	}
	return "YES"
}

func main() {
	file, _ := os.Open("assets/balancedBreaker.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	firstLine := 0
	res := make([]string, 0)
	for scanner.Scan() {
		s := scanner.Text()
		if firstLine == 0 {
			firstLine++
			continue
		}
		res = append(res, isBalanced(s))
		firstLine++
	}

	file, _ = os.Open("assets/expectBalancedBreaker.txt")
	defer file.Close()
	scanner = bufio.NewScanner(file)
	resExpect := make([]string, 0)
	for scanner.Scan() {
		v := scanner.Text()
		resExpect = append(resExpect, v)
	}

	for i, v := range res {
		if v != resExpect[i] {
			fmt.Printf("dont expect for line %d: %s\n", i, v)
		}
	}
}
