package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	considers := make(map[int]int)
	for i, num := range nums {
		consider := target - num

		if v, ok := considers[consider]; ok {
			return []int{v, i}
		}
		considers[num] = i
	}
	return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	fmt.Println(twoSum(nums, target))
}
