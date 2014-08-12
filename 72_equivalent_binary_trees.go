package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

const TREE_SIZE = 10

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	a1 := readTree(t1)
	a2 := readTree(t2)

	return compareSlices(a1, a2)
}

func readTree(t *tree.Tree) []int {
	ch := make(chan int)
	go Walk(t, ch)

	contents := make([]int, TREE_SIZE)
	for i := 0; i < TREE_SIZE; i++ {
		contents[i] = <-ch
	}

	return contents
}

func compareSlices(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}

	for i, value := range a1 {
		if value != a2[i] {
			return false
		}
	}

	return true
}

func main() {
	t1 := tree.New(1)
	t1prime := tree.New(1)
	t2 := tree.New(2)

	if Same(t1, t1prime) {
		fmt.Println("Success: 1-based trees are equal as expected")
	} else {
		fmt.Println("Error: 1-based trees are not equal, but should be")
	}

	if Same(t1, t2) {
		fmt.Println("Error: 1- and 2-based trees are equal, but shouldn't be")
	} else {
		fmt.Println("Success: 1- and 2-based trees are not equal")
	}
}
