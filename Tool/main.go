package main

import "fmt"

func main() {
	ints := []int{1, 11, 19}
	insert(&ints, map[int]int{1: 8, 7: 7})
	fmt.Println(ints[0:])
}
