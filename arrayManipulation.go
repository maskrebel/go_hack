package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int64, n+1)
	var maxVal int64
	var currentSum int64

	for _, q := range queries {
		start, end, val := int64(q[0]), int64(q[1]), int64(q[2])

		arr[start-1] += val
		if end < int64(n) {
			arr[end] -= val
		}
	}

	for _, v := range arr {
		currentSum += v
		if currentSum > maxVal {
			maxVal = currentSum
		}
	}

	return maxVal
}

func main() {
	file, _ := os.Open("assets/arrayManipulation.txt")
	defer file.Close()

	queries := make([][]int32, 0)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstLine := scanner.Text()
	splitFirstLine := strings.Split(firstLine, " ")
	n, _ := strconv.Atoi(splitFirstLine[0])
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, " ")
		query := make([]int32, 0)
		for _, v := range lineArr {
			q, _ := strconv.Atoi(v)
			query = append(query, int32(q))
		}
		queries = append(queries, query)
	}
	//n := int32(5)
	//queries := [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}
	fmt.Println(arrayManipulation(int32(n), queries))
}
