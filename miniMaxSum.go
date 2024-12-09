package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	firstVal := int64(arr[0])
	lastVal := int64(arr[len(arr)-1])

	total := int64(0)
	for _, v := range arr {
		total += int64(v)
	}

	fmt.Println(total-lastVal, total-firstVal)
}

func main() {
	arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	miniMaxSum(arr)

}
