package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int32

// Implementing heap.Interface
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int32))
}

func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func cookies(k int32, A []int32) int32 {
	// Create a min-heap
	h := &IntHeap{}
	heap.Init(h)

	// Push all elements of A into the heap
	for _, a := range A {
		heap.Push(h, a)
	}

	opc := int32(0)
	// While the smallest element is less than k, keep mixing cookies
	for h.Len() > 1 && (*h)[0] < k {
		// Pop the two smallest elements
		a := heap.Pop(h).(int32)
		b := heap.Pop(h).(int32)

		// Push the new cookie back to the heap
		heap.Push(h, a+2*b)

		opc++
	}

	// If the smallest element is less than k, return -1
	if (*h)[0] < k {
		return -1
	}
	return opc
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	idx := 0
	A := make([]int32, 0)
	k := int32(0)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		if idx == 0 {
			tmpK, _ := strconv.Atoi(lineSplit[1])
			k = int32(tmpK)
			idx++
			continue
		}
		for _, v := range lineSplit {
			val, _ := strconv.Atoi(v)
			A = append(A, int32(val))
		}

	}
	fmt.Println(cookies(k, A))
}
