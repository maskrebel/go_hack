package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func flippingMatrix(matrix [][]int32) int32 {
	//fmt.Println(max(matrix[0][0], matrix[0][3], matrix[3][0], matrix[3][3])) // 1 4 13 16 max -> 16
	//fmt.Println(max(matrix[0][1], matrix[0][2], matrix[3][1], matrix[3][2])) // 2 5 14 15 max -> 15
	//fmt.Println(max(matrix[1][0], matrix[1][3], matrix[2][0], matrix[2][3])) // 5 8 5 12 max -> 12
	//fmt.Println(max(matrix[1][1], matrix[1][2], matrix[2][1], matrix[2][2])) // 5 6 10 11 max -> 11
	// sum 16+15+12+11 = 54
	n := len(matrix) / 2
	maximum := int32(0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val1 := matrix[i][j]
			val2 := matrix[i][len(matrix)-1-j]
			val3 := matrix[len(matrix)-1-i][j]
			val4 := matrix[len(matrix)-1-i][len(matrix)-1-j]
			maximum += max(val1, val2, val3, val4)
			fmt.Println(max(val1, val2, val3, val4))
		}
	}
	return maximum
}

func main() {
	//sample input
	//1 2 3 4
	//5 6 7 8
	//9 10 11 12
	//13 14 15 16

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]int32
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, " ")
		newArr := make([]int32, 0)
		for _, val := range lineSlice {
			convVal, _ := strconv.Atoi(val)
			newArr = append(newArr, int32(convVal))
		}
		matrix = append(matrix, newArr)
	}
	fmt.Println(flippingMatrix(matrix))
}
