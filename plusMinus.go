package main

import "fmt"

func plusMinus(arr []int32) {
	plus := 0
	minus := 0
	zero := 0
	for _, v := range arr {
		if v > 0 {
			plus++
		} else if v < 0 {
			minus++
		} else {
			zero++
		}
	}

	fmt.Printf("%.6f\n", float32(plus)/float32(len(arr)))
	fmt.Printf("%.6f\n", float32(minus)/float32(len(arr)))
	fmt.Printf("%.6f\n", float32(zero)/float32(len(arr)))
}

func main() {
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})
}
