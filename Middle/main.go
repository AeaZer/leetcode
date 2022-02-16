package main

import "fmt"

func main() {
	ints := [][]int{{0, 1}, {3, 4}, {2, 4}, {7, 9}, {0, 3}}
	i := merge(ints)
	fmt.Println(i)
}
