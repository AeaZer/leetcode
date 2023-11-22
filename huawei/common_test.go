package huawei

import (
	"strconv"
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

func TestPerfectCity(t *testing.T) {
	ans := ipStrToInt("1.1.1.15")
	t.Log(strconv.FormatInt(ans, 2))
}

func TestActualCost(t *testing.T) {
	t.Log(actualCost(225))
	t.Log(actualCost(15))
}
