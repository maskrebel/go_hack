package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countingSort(arr []int32) []int32 {
	newArr := make([]int32, 100)
	for _, v := range arr {
		newArr[v] += 1
	}

	return newArr
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr []int32
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		for _, v := range lineSplit {
			num, _ := strconv.Atoi(v)
			arr = append(arr, int32(num))
		}
	}
	fmt.Println(countingSort(arr))
}
