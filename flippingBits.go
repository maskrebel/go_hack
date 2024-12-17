package main

import (
	"fmt"
	"strconv"
)

func flip(s string) (res string) {
	for _, v := range s {
		if string(v) == "1" {
			res = res + "0"
		} else {
			res = res + "1"
		}
	}
	return res
}

func make32Byte(x int) (res string) {
	for i := 0; i < x; i++ {
		res = "1" + res
	}
	return res
}

func flippingBits(n int64) int64 {
	bitStr := strconv.FormatInt(n, 2)
	bitStr = flip(bitStr)
	bitStr = make32Byte(32-len(bitStr)) + bitStr
	decimalStr, _ := strconv.ParseInt(bitStr, 2, 64)
	return decimalStr
}

func main() {
	fmt.Println(flippingBits(9))

}
