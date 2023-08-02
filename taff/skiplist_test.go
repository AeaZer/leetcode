package taff

import (
	"fmt"
	"testing"
)

func TestSkip(t *testing.T) {
	sl := newSkipList()

	sl.insert(3)
	sl.insert(6)
	sl.insert(2)
	sl.insert(8)
	sl.insert(1)
	sl.insert(5)

	sl.print()

	fmt.Println("Search 5:", sl.search(5))
	fmt.Println("Search 4:", sl.search(4))

	sl.delete(6)

	sl.print()
}

func TestRange(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		ints = append(ints, v)
	}
	fmt.Println(ints)
}
