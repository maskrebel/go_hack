package main

import "fmt"

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
	q := []int32{2, 1, 5, 3, 4}
	minimumBribes(q)
}
