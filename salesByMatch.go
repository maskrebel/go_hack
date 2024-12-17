package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sockMerchant(n int32, ar []int32) int32 {
	mapInt := make(map[int32]int32)
	for _, v := range ar {
		mapInt[v]++
	}

	count := int32(0)
	for _, t := range mapInt {
		count += t / 2
	}
	return count
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lineN := scanner.Text()
	n, _ := strconv.Atoi(lineN)
	ar := make([]int32, 0)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	for _, v := range line {
		num, _ := strconv.Atoi(v)
		ar = append(ar, int32(num))
	}
	fmt.Println(sockMerchant(int32(n), ar))

}
