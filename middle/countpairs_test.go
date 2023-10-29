package main

import "testing"

func TestCountPairs(t *testing.T) {
	countPairs(7, [][]int{{0, 15}, {1, 14}, {2, 11}, {4, 3}, {5, 15}, {8, 2}, {14, 12}})
}
