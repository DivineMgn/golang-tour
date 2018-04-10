// https://go-tour-ru-ru.appspot.com/concurrency/8

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	channel := make(chan int)
	firstTree, secondTree := tree.New(1), tree.New(3)

	go Walk(firstTree, channel)
	for value := range channel {
		fmt.Println(value)
	}
	//fmt.Println("fitst tree:\n", firstTree)
	//fmt.Println("second tree:\n", secondTree)
	fmt.Println(Same(firstTree, firstTree))
	fmt.Println(Same(firstTree, secondTree))
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(someTree *tree.Tree, channel chan<- int) {
	defer close(channel)
	// thread unsafe collection
	queue := []*tree.Tree{someTree}

	for len(queue) > 0 {
		curTree := queue[0]
		queue = queue[1:]

		if curTree == nil {
			continue
		}

		channel <- curTree.Value

		if curTree.Left != nil {
			queue = append(queue, curTree.Left)
		}

		if curTree.Right != nil {
			queue = append(queue, curTree.Right)
		}
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(tree1, tree2 *tree.Tree) bool {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go Walk(tree1, channel1)
	go Walk(tree2, channel2)

	for value1 := range channel1 {
		if value1 != <-channel2 {
			return false
		}
	}

	return true
}
