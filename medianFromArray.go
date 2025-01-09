package main

import (
	"fmt"
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := append(nums1, nums2...)
	sort.Ints(arr)

	var med int
	if len(arr)%2 != 0 {
		med = len(arr) / 2
		return math.Round(float64(arr[med]) * 1e6)
	} else {
		med = len(arr) / 2
		res := arr[med+1] + arr[med+2]
		return math.Round(float64(res/2) * 1e6)
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))

}
