package huawei

import (
	"testing"
)

func TestFindMaxParallelism(t *testing.T) {
	ans := findMaxParallelism([][3]int{{1, 10, 3}, {2, 20, 2}, {3, 30, 1}})
	t.Log(ans)

	ans = findMaxParallelism([][3]int{{3, 5, 2}, {6, 7, 3}})
	t.Log(ans)
}

func TestWordFlashback(t *testing.T) {
	ans := wordFlashback("hello im is areazer, do you have any problem? haha")
	t.Log(ans)
}
