package hot100

func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		m[v] = true
	}
	passed := make(map[int]bool, len(m))
	var longest int
	for v := range m {
		if passed[v] {
			continue
		}
		current := 1
		for m[v+1] {
			passed[v] = true
			current++
			v++
		}
		if current > longest {
			longest = current
		}
	}
	return longest
}
