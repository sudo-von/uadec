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
			Nodes: []string{"2", "3", "4", "5"},
		},
		"5": {
			Value: [][]string{
				{"1", "4", "3"}, {"7", "6", "0"}, {"5", "8", "2"},
			},
			Nodes: []string{"12", "13"},
		},
		"4": {
			Value: [][]string{
				{"1", "4", "3"}, {"7", "8", "6"}, {"5", "0", "2"},
			},
			Nodes: []string{"10", "11"},
		},
		"3": {
			Value: [][]string{
				{"1", "4", "3"}, {"0", "7", "6"}, {"5", "8", "2"},
			},
			Nodes: []string{"8", "9"},
		},
		"2": {
			Value: [][]string{
				{"1", "0", "3"}, {"7", "4", "6"}, {"5", "8", "2"},
			},
			Nodes: []string{"6", "7"},
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

	fmt.Println("Using queue [FIFO]")
	fmt.Println("Showing queue...")

	queue := structures.NewQueue()
	var path []string

	queue.Push(root)
	fmt.Println(queue)

	for queue.Count() != 0 {

		data := *queue.Pop()
		path = append(path, data)

		for _, key := range binaryTree[data].Nodes {
			queue.Push(key)
		}
		fmt.Println(queue)
	}

	fmt.Println("Queue is empty...")
	fmt.Printf("Path: %s \n", path)

	fmt.Println("Printing path but instead of nodes i will print the data...")
	for _, v := range path {
		fmt.Println(binaryTree[v].Value)
	}

}
