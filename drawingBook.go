package main

import "fmt"

func pageCount(n int32, p int32) int32 {
	totalPage := (n / 2) + 1

	res := int32(0)
	for i := int32(0); i < totalPage; i++ {
		if i+i == p || i+i+1 == p {
			res = i
			break
		}

		if n%2 == 0 && (n-i-i == p || n-i-i+1 == p) {
			res = i
			break
		} else if n%2 != 0 && (n-i-i == p || n-i-i-1 == p) {
			res = i
			break
		}
	}
	return res
}

func main() {
	fmt.Println(pageCount(6, 2) == 1)
	fmt.Println(pageCount(5, 4) == 0)
	fmt.Println(pageCount(6, 5) == 1)
}
