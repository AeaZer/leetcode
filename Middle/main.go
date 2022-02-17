package main

import "fmt"

func main() {
	ints := [][]int{{2, 3}, {2, 2}, {3, 4}, {5, 5}, {2, 4}}
	i := merge(ints)
	fmt.Println(i)
}
