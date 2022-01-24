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
			Nodes: []string{"7", "6", "5", "4", "3", "2"},
		},
		"2": {
			Value: [][]string{
				{"a", "b"}, {"c"},
			},
			Nodes: []string{"9", "8"},
		},
		"3": {
			Value: [][]string{
				{"b"}, {"a", "c"},
			},
			Nodes: []string{"11", "10"},
		},
		"4": {
			Value: [][]string{
				{"b", "a"}, {"c"},
			},
			Nodes: []string{"13", "12"},
		},
		"5": {
			Value: [][]string{
				{"a"}, {"b", "c"},
			},
			Nodes: []string{"15", "14"},
		},
		"6": {
			Value: [][]string{
				{"c", "a"}, {"b"},
			},
			Nodes: []string{"17", "16"},
		},
		"7": {
			Value: [][]string{
				{"a"}, {"c", "b"},
			},
			Nodes: []string{"19", "18"},
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
