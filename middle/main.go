package main

import "fmt"

func main() {
	value, ok := newTree().searchKth(2)
	fmt.Println(value)
	fmt.Println(ok)
}
