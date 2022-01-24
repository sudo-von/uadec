package main

import (
	"fmt"
	"uadec/breadth-firsth-search/structures"
)

func main() {

	binaryTree := map[string]structures.Node{
		"1": {
			Value: [][]string{
				{"1", "4", "3"}, {"7", "0", "6"}, {"5", "8", "2"},
			},
			Nodes: []string{"5", "4", "3", "2"},
		},
		"5": {
			Value: [][]string{
				{"1", "4", "3"}, {"7", "6", "0"}, {"5", "8", "2"},
			},
			Nodes: []string{"13", "12"},
		},
		"4": {
			Value: [][]string{
				{"1", "4", "3"}, {"7", "8", "6"}, {"5", "0", "2"},
			},
			Nodes: []string{"11", "10"},
		},
		"3": {
			Value: [][]string{
				{"1", "4", "3"}, {"0", "7", "6"}, {"5", "8", "2"},
			},
			Nodes: []string{"9", "8"},
		},
		"2": {
			Value: [][]string{
				{"1", "0", "3"}, {"7", "4", "6"}, {"5", "8", "2"},
			},
			Nodes: []string{"7", "6"},
		},
		"11": {
			Value: [][]string{
				{"1", "4", "3"}, {"7", "8", "6"}, {"5", "2", "0"},
			},
			Nodes: []string{},
		},
		"10": {
			Value: [][]string{
				{"1", "4", "3"}, {"7", "8", "6"}, {"0", "5", "2"},
			},
			Nodes: []string{},
		},
		"9": {
			Value: [][]string{
				{"1", "4", "3"}, {"5", "7", "6"}, {"0", "8", "2"},
			},
			Nodes: []string{},
		},
		"8": {
			Value: [][]string{
				{"0", "4", "3"}, {"1", "7", "6"}, {"5", "8", "2"},
			},
			Nodes: []string{},
		},
		"7": {
			Value: [][]string{
				{"1", "3", "0"}, {"7", "4", "6"}, {"5", "8", "2"},
			},
			Nodes: []string{},
		},
		"6": {
			Value: [][]string{
				{"0", "1", "3"}, {"7", "4", "6"}, {"5", "8", "2"},
			},
			Nodes: []string{},
		},
		"13": {
			Value: [][]string{
				{"1", "4", "3"}, {"7", "6", "2"}, {"5", "8", "0"},
			},
			Nodes: []string{},
		},
		"12": {
			Value: [][]string{
				{"1", "4", "0"}, {"7", "6", "3"}, {"5", "8", "2"},
			},
			Nodes: []string{},
		},
	}

	breadthFirstSearch(binaryTree, "1")

}

func breadthFirstSearch(binaryTree map[string]structures.Node, root string) {

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

		for _, key := range binaryTree[data].Nodes {
			stack.Push(key)
		}
		fmt.Println(stack)
	}

	fmt.Println("Stack is empty...")
	fmt.Printf("Path: %s \n", path)

	fmt.Println("Printing path but instead of nodes i will print the data...")
	for _, v := range path {
		fmt.Println(binaryTree[v].Value)
	}

}
