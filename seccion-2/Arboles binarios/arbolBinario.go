package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk2(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk2(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk2(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk2(t1, ch1)
	go Walk2(t2, ch2)

	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk2(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Estos arboles binarios son equivalentes", Same(tree.New(1), tree.New(1)))
	fmt.Println("Estos arboles binarios no son equivalentes", Same(tree.New(1), tree.New(2)))
}
