package main

import "fmt"

func sumArray(arr []int32) int32 {
	total := int32(0)
	for _, v := range arr {
		total += v
	}
	return total
}

func birthday(s []int32, d int32, m int32) int32 {
	count := int32(0)
	for i := int32(0); i < (int32(len(s)) - m + 1); i++ {
		amount := sumArray(s[i : i+m])
		if amount == d {
			count++
		}
	}

	return count
}

func main() {
	s := []int32{1, 2, 1, 3, 2}
	d := int32(3)
	m := int32(2)
	fmt.Println(birthday(s, d, m))
}
