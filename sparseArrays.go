package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateSparseArrays(a string, str []string) (total int32) {
	for _, s := range str {
		if a == s {
			total++
		}
	}
	return total
}

func matchingStrings(strings []string, queries []string) []int32 {
	ArrRes := make([]int32, len(queries))
	for i, q := range queries {
		ArrRes[i] = calculateSparseArrays(q, strings)
	}
	return ArrRes
}

func main() {
	file, _ := os.Open("assets/sparseArrays.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	strArr := make([]string, 0)
	queryArr := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		strCount, _ := strconv.Atoi(line)
		for i := 0; i < strCount; i++ {
			scanner.Scan()
			strArr = append(strArr, scanner.Text())
		}
		scanner.Scan()
		lines := scanner.Text()
		queryCount, _ := strconv.Atoi(lines)
		for i := 0; i < queryCount; i++ {
			scanner.Scan()
			queryArr = append(queryArr, scanner.Text())
		}
		break
	}
	fmt.Println(matchingStrings(strArr, queryArr))
}
