package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func gridChallenge(grid []string) string {
	for i := 0; i < len(grid[0]); i++ { // grid[0] sample for indexes
		greaterValidate := int32(0)
		for _, row := range grid {
			runes := []rune(row)
			sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
			if greaterValidate > runes[i] {
				return "NO"
			} else {
				greaterValidate = runes[i]
			}
		}
	}
	return "YES"
}

func main() {
	// sample payload
	// sort every row, and make new array by horizontal
	//////for//////////////to///////////
	//	eabcd			afkpu
	//	fghij			bglqv
	//	olkmn			chmrw
	//	trpqs			dinsx
	//	xywuv			ejoty

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	fmt.Println(gridChallenge(grid))
}
