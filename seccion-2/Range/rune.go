package main

import "fmt"

func main() {
	var string = "GeeksforGeeks"
	for i, item := range string {
		fmt.Printf("odd[%d] = %d \n", i, item)
	}
}
