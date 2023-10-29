package _0231029

import "testing"

func TestMinSum(t *testing.T) {
	minSum([]int{20, 0, 18, 11, 0, 0, 0, 0, 0, 0, 17, 28, 0, 11, 10, 0, 0, 15, 29}, []int{16, 9, 25, 16, 1, 9, 20, 28, 8, 0, 1, 0, 1, 27})
}

func TestMinSum2(t *testing.T) {
	minSum([]int{3, 2, 0, 1, 0}, []int{6, 5, 0})
}
