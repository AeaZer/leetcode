package dp

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	f1, f2, i := cost[0], cost[1], 2
	if n == 2 {
		return f2
	}
	for i < n {
		f1, f2 = f2, min(f1, f2)+cost[i]
		i++
	}
	return min(f1, f2)
}
