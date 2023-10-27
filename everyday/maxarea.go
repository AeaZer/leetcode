package everyday

import (
	"math"
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = append(horizontalCuts, []int{0, h}...)
	verticalCuts = append(verticalCuts, []int{0, w}...)
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	hMax, vMax := math.MinInt, math.MinInt
	for i := 1; i < len(horizontalCuts); i++ {
		gap := horizontalCuts[i] - horizontalCuts[i-1]
		if gap > hMax {
			hMax = gap
		}
	}
	for i := 1; i < len(verticalCuts); i++ {
		gap := verticalCuts[i] - verticalCuts[i-1]
		if gap > vMax {
			vMax = gap
		}
	}
	return (hMax * vMax) % (1e9 + 7)
}
