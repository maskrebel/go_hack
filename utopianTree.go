package main

import "fmt"

func utopianTree(n int32) int32 {
	mapTreeTall := func() int32 {
		highCount := n + 1
		total := int32(0)
		for i := int32(0); i < highCount; i++ {
			if i%2 == 1 {
				total += total
			} else {
				total++
			}
		}
		return total
	}

	return mapTreeTall()
}

func main() {
	//	Period  Height
	//	0          1
	//	1          2
	//	2          3
	//	3          6
	//	4          7
	//	5          14

	fmt.Println(utopianTree(int32(0)))
	fmt.Println(utopianTree(int32(1)))
	fmt.Println(utopianTree(int32(4)))
	fmt.Println(utopianTree(int32(3)))
	fmt.Println(utopianTree(int32(10)))
}
