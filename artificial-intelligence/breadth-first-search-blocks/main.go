package main

import (
	"fmt"
	"uadec/breadth-firsth-search/structures"
)

func main() {

	binaryTree := map[string]structures.Node{
		"1": {
			Value: [][]string{
				{"a"}, {"b"}, {"c"},
			},
			Nodes: []string{"2", "3", "4", "5", "6", "7"},
		},
		"2": {
			Value: [][]string{
				{"a", "b"}, {"c"},
			},
			Nodes: []string{"8", "9"},
		},
		"3": {
			Value: [][]string{
				{"b"}, {"a", "c"},
			},
			Nodes: []string{"10", "11"},
		},
		"4": {
			Value: [][]string{
				{"b", "a"}, {"c"},
			},
			Nodes: []string{"12", "13"},
		},
		"5": {
			Value: [][]string{
				{"a"}, {"b", "c"},
			},
			Nodes: []string{"14", "15"},
		},
		"6": {
			Value: [][]string{
				{"c", "a"}, {"b"},
			},
			Nodes: []string{"16", "17"},
		},
		"7": {
			Value: [][]string{
				{"a"}, {"c", "b"},
			},
			Nodes: []string{"18", "19"},
		},
		"8": {
			Value: [][]string{
				{"c", "a", "b"},
			},
			Nodes: []string{},
		},
		"9": {
			Value: [][]string{
				{"b"}, {"a", "c"},
			},
			Nodes: []string{},
		},
		"10": {
			Value: [][]string{
				{"b", "a", "c"},
			},
			Nodes: []string{},
		},
		"11": {
			Value: [][]string{
				{"a", "b"}, {"c"},
			},
			Nodes: []string{},
		},
		"12": {
			Value: [][]string{
				{"c", "b", "a"},
			},
			Nodes: []string{},
		},
		"13": {
			Value: [][]string{
				{"a"}, {"b", "c"},
			},
			Nodes: []string{},
		},
		"14": {
			Value: [][]string{
				{"a", "b", "c"},
			},
			Nodes: []string{},
		},
		"15": {
			Value: [][]string{
				{"b", "a"}, {"c"},
			},
			Nodes: []string{},
		},
		"16": {
			Value: [][]string{
				{"b", "c", "a"},
			},
			Nodes: []string{},
		},
		"17": {
			Value: [][]string{
				{"a"}, {"c", "b"},
			},
			Nodes: []string{},
		},
		"18": {
			Value: [][]string{
				{"a", "c", "b"},
			},
			Nodes: []string{},
		},
		"19": {
			Value: [][]string{
				{"c", "a"}, {"b"},
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
