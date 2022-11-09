package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		fmt.Println("---------IZQUIERDA---------")
		fmt.Println("valor del nodo ", t.Value)
		fmt.Println("valor del lado izquierdo ", t.Left)
		fmt.Println("valor del lado derecho ", t.Right)
		fmt.Println("--------------------")
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		fmt.Println("---------DERECHA---------")
		fmt.Println("valor del nodo ", t.Value)
		fmt.Println("valor del lado izquierdo ", t.Left)
		fmt.Println("valor del lado derecho ", t.Right)
		fmt.Println("--------------------")
		Walk(t.Right, ch)
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(2), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
