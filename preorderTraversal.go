package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func insertNodes(node *Node, value int) {
	if node.Value > value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
			return
		}
		insertNodes(node.Left, value)
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value}
			return
		}
		insertNodes(node.Right, value)
	}

}

func contructTree(nodes []int) *Node {
	root := &Node{Value: nodes[0]}

	for i := 1; i < len(nodes); i++ {
		insertNodes(root, nodes[i])
	}
	return root
}

func printTree(node *Node) {
	fmt.Printf("%d ", node.Value)
	if node.Left != nil {
		printTree(node.Left)
	}

	if node.Right != nil {
		printTree(node.Right)
	}
}

func main() {
	// sample input
	// 6
	// 1 2 5 3 6 4
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	total, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	arr := make([]int, total)
	scanner.Scan()
	tmpArr := strings.Split(scanner.Text(), " ")
	for i, v := range tmpArr {
		val, _ := strconv.Atoi(v)
		arr[i] = val
	}

	root := contructTree(arr)
	printTree(root)

}
