package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//type MapWordInt map[string]int

func designerPdfViewer(h []int32, word string) int32 {
	MapWordInt := func(h []int32) map[string]int32 {
		if len(h) != 26 {
			h = make([]int32, 26)
		}
		resMap := make(map[string]int32)
		for i, j := int32(97), 0; i <= int32(122); i, j = i+1, j+1 {
			resMap[string(i)] = h[j]
		}
		return resMap
	}

	wordInt := MapWordInt(h)
	tallLetter := int32(0)
	for _, a := range word {
		tall := wordInt[string(a)]
		if tall > tallLetter {
			tallLetter = tall
		}
	}

	countWord := int32(len(word))
	return tallLetter * countWord
}

func main() {
	// sample input
	//	1 3 1 3 1 4 1 3 2 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 7
	//	zaba

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	arr := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		arr = append(arr, line)
	}

	rawHighs := arr[0]
	splitHighs := strings.Split(rawHighs, " ")
	highs := make([]int32, len(splitHighs))
	for i, s := range splitHighs {
		val, _ := strconv.Atoi(s)
		highs[i] = int32(val)
	}
	word := arr[1]
	fmt.Println(designerPdfViewer(highs, word))
}
