package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func arrayToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}

func reverseIntToString(l *ListNode) string {
	var strInt string
	for l.Next != nil {
		strInt = strconv.Itoa(l.Val) + strInt
		l = l.Next
	}
	strInt = strconv.Itoa(l.Val) + strInt
	return strInt
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Str := reverseIntToString(l1)
	l2Str := reverseIntToString(l2)
	if len(l1Str) > len(l2Str) {
		l2Str = strings.Repeat("0", len(l1Str)-len(l2Str)) + l2Str
	} else {
		l1Str = strings.Repeat("0", len(l2Str)-len(l1Str)) + l1Str
	}

	l1Arr := strings.Split(l1Str, "")
	l2Arr := strings.Split(l2Str, "")

	var arr []int
	var extend bool
	for i := len(l1Str) - 1; i >= 0; i-- {
		var a, _ = strconv.Atoi(l1Arr[i])
		var b, _ = strconv.Atoi(l2Arr[i])
		sum := a + b
		if extend {
			sum++
			extend = false
		}

		if sum >= 10 {
			sum -= 10
			extend = true
		}

		arr = append(arr, sum)
	}

	if extend {
		arr = append(arr, 1)
	}
	fmt.Println(arr)
	return arrayToList(arr)
}

func main() {
	//arrL1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	arrL1 := []int{2, 4, 3}
	arrL2 := []int{5, 6, 4}
	l1 := arrayToList(arrL1)
	l2 := arrayToList(arrL2)

	res := addTwoNumbers(l1, l2)
	fmt.Println(res)
}
