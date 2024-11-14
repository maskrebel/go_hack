package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func caesarCipher(s string, k int32) string {
	res := ""
	k %= 26

	for _, v := range s {
		if unicode.IsLetter(v) {
			beforeTransform := unicode.IsUpper(v)
			v += k
			for {
				if !unicode.IsLetter(v) {
					v -= 26
				} else {
					break
				}
			}
			afterTransform := unicode.IsUpper(v)

			if beforeTransform != afterTransform {
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
