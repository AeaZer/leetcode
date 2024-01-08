package everyday

func numberOfBoomerangs(points [][]int) int {
	var ans int
	for _, point := range points {
		distMap := make(map[int]int)
		for j := 0; j < len(points); j++ {
			distMap[getDistance([2][]int{point, points[j]})]++
		}
		for _, count := range distMap {
			if count > 1 {
				ans += combineCn2(count)
			}
		}
	}
	return ans * 2
}

func getDistance(points [2][]int) int {
	r, c := points[0][0]-points[1][0], points[0][1]-points[1][1]
	return r*r + c*c
}

func combineCn2(n int) int {
	return n * (n - 1) / 2
}
