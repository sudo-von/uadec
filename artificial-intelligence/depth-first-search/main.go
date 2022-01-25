package main

import (
	"fmt"
	"uadec/depth-firsth-search/structures"
)

func main() {

	binaryTree := map[string][]string{
		"1": {"3", "2"},
		"3": {"7", "6"},
		"2": {"5", "4"},
		"5": {"9", "8"},
		"7": {"10"},
		"8": {"12", "11"},
	}

	depthFirstSearch(binaryTree, "1")

}

func depthFirstSearch(binaryTree map[string][]string, root string) {

	if _, found := binaryTree[root]; !found {
		fmt.Println("Root not found")
		return
	}

	fmt.Println("Using stack [LIFO]")
	fmt.Println("Showing stack...")

	stack := structures.NewStack()
	var path []string

	stack.Push(root)
	fmt.Println(stack)

	for stack.Count() != 0 {

		data := *stack.Pop()
		path = append(path, data)

		for _, key := range binaryTree[data] {
			stack.Push(key)
		}
		fmt.Println(stack)
	}

	fmt.Println("Stack is empty...")
	fmt.Printf("Path: %s", path)

}
