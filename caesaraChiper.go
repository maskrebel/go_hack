package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func isAlphabet(c rune) bool {
	alphabets := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S',
		'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
		'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	for _, alphabet := range alphabets {
		if c == alphabet {
			return true
		}
	}
	return false
}

func validateUpperLower(v rune) string {
	res := ""
	if unicode.IsUpper(v) {
		res = "upper"
	} else {
		res = "lower"
	}

	return res
}

func caesarCipher(s string, k int32) string {
	res := ""
	for {
		if k > 26 {
			k -= 26
		} else {
			break
		}
	}

	for _, v := range s {
		if isAlphabet(v) {
			validate := validateUpperLower(v)
			v += k
			for {
				if !isAlphabet(v) {
					v -= 26
				} else {
					break
				}
			}

			afterVal := validateUpperLower(v)

			if validate != afterVal {
				v -= 26
			}
		}
		res += string(v)
	}
	return res
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr []string
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}
	str := arr[0]
	trim, _ := strconv.Atoi(arr[1])
	fmt.Println(str)
	fmt.Println(caesarCipher(str, int32(trim)))
	val := "!w-bL`-yX!.G`mVKmFlX/MaCyyvSS!CSwts.!g/`He`eJk1DGZBa`eBLdyu`hZD`vV-jZDm.LCBSre..-!.!dmv!-E"
	fmt.Println(val)
	fmt.Println(val == caesarCipher(str, int32(trim)))
}
