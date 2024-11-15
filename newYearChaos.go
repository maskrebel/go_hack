package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minimumBribes(q []int32) {
	bribes := 0
	for i := len(q) - 1; i >= 0; i-- {
		if q[i]-int32(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}

		for j := max(0, int(q[i]-2)); j < i; j++ {
			if q[j] > q[i] {
				bribes++
			}
		}
	}
	fmt.Println(bribes)
	return
}

func main() {
	file, _ := os.Open("assets/newYearChaos.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	arr := make([][]int32, 0)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		i++
		if i == 1 || i%2 == 0 {
			continue
		}
		splitLines := strings.Split(line, " ")
		sArr := make([]int32, 0)
		for _, v := range splitLines {
			newVal, _ := strconv.Atoi(v)
			sArr = append(sArr, int32(newVal))
		}
		arr = append(arr, sArr)
	}

	for _, q := range arr {
		minimumBribes(q)
	}
}
