package competitions

import (
	"sort"
)

type beautifulPairs struct {
	fistIndex int
	lastIndex int
	str       string
}

func shortestBeautifulSubstring(s string, k int) string {
	pairs := make([]beautifulPairs, 0)
	sl := len(s)
	for i := 0; i < sl-k+1; i++ {
		c := k
		j := i
		for j < sl && c > 0 {
			if s[j] == '1' {
				c--
			}
			j++
		}
		if j <= sl && c == 0 {
			pairs = append(pairs, beautifulPairs{
				fistIndex: i,
				lastIndex: j,
				str:       s[i:j],
			})
		}
	}
	if len(pairs) == 0 {
		return ""
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].fistIndex-pairs[i].lastIndex != pairs[j].fistIndex-pairs[j].lastIndex {
			return pairs[i].fistIndex-pairs[i].lastIndex > pairs[j].fistIndex-pairs[j].lastIndex
		}
		return pairs[i].str < pairs[j].str
	})
	return s[pairs[0].fistIndex:pairs[0].lastIndex]
}
