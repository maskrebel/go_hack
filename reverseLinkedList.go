package main

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode) {
	for node != nil {
		fmt.Println(node.data)

		node = node.next
	}
}

func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	arr := make([]int32, 0)
	newLlist := SinglyLinkedList{}
	for llist != nil {
		arr = append(arr, llist.data)
		llist = llist.next
	}
	//sort.Slice(arr, func(i, j int) bool { return arr[i] > arr[j] })
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	for _, v := range arr {
		newLlist.insertNodeIntoSinglyLinkedList(v)
	}
	return newLlist.head
}

func main() {
	llist := SinglyLinkedList{}
	for i := int32(1); i <= 5; i++ {
		llist.insertNodeIntoSinglyLinkedList(i)
	}
	llist1 := reverse(llist.head)
	printSinglyLinkedList(llist1)
}
