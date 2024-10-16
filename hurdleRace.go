package main

import (
	"fmt"
	"sort"
)

func hurdleRace(k int32, height []int32) int32 {
	// k for maximum value
	// you can use sort.Ints if type is (int) not int32 or others
	sort.Slice(height, func(i, j int) bool { return height[i] < height[j] })
	maxHeight := height[len(height)-1]
	res := maxHeight - k
	if res < 0 {
		res = 0
	}

	return res

}

func main() {
	arr := []int32{5, 1, 2, 3, 4}
	k := int32(7)
	res := hurdleRace(k, arr)
	fmt.Println(res)
}
