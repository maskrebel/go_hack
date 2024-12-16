package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// search unique number
// ex: [1 2 3 4 3 2 1] and unique is 4
func lonelyInteger(a []int32) int32 {
	mapArr := make(map[int32]struct{})
	for _, v := range a {
		if _, ok := mapArr[v]; ok {
			delete(mapArr, v)
		} else {
			mapArr[v] = struct{}{}
		}
	}

	unique := int32(0)
	for i := range mapArr {
		unique = i
	}

	return unique
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
	fmt.Println(lonelyInteger(arr))
}
