package main

import "fmt"

func truckTour(petrolpumps [][]int32) int32 {
	// Write your code here
	balance, deficit, start := int32(1), int32(1), int32(1)
	for i := int32(0); i < int32(len(petrolpumps)); i++ {
		balance += petrolpumps[i][0] - petrolpumps[i][1]
		balance += 1
		if balance < 0 {
			deficit += balance
			start = i + 1
			balance = 0
		}
	}

	if balance+deficit >= 0 {
		return start
	}

	return -1
}

func main() {
	p := [][]int32{{1, 5}, {10, 3}, {3, 4}}
	fmt.Println(truckTour(p))
}
