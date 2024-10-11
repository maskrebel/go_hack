package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	lDiagonal := int32(0)
	rDiagonal := int32(0)
	idxL := 0
	idxR := len(arr) - 1
	for _, val := range arr {
		lDiagonal += val[idxL]
		rDiagonal += val[idxR]
		idxL++
		idxR--
	}
	total := lDiagonal - rDiagonal
	if total < 0 {
		total *= -1
	}
	return total
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr [][]int32
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		var arrSingle []int32
		for _, v := range lineSplit {
			num, _ := strconv.Atoi(v)
			arrSingle = append(arrSingle, int32(num))
		}
		arr = append(arr, arrSingle)
	}
	fmt.Println(diagonalDifference(arr))
}
