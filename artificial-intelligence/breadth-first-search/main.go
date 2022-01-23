package main

import (
	"fmt"
	"uadec/breadth-firsth-search/structures"
)

func main() {

	binaryTree := map[string][]string{
		"1": {"2", "3"},
		"3": {"6", "7"},
		"2": {"4", "5"},
		"5": {"8", "9"},
		"7": {"10"},
		"8": {"11", "12"},
	}

	breadthFirstSearch(binaryTree, "1")

}

func breadthFirstSearch(binaryTree map[string][]string, root string) {

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

		for _, key := range binaryTree[data] {
			queue.Push(key)
		}
		fmt.Println(queue)
	}
	fmt.Println("Queue is empty...")
	fmt.Printf("Path: %s", path)
}
