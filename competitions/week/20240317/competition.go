package _0240717

import (
	"sort"
	"strings"

	"golang.org/x/text/unicode/rangetable"
)

func isSubstringPresent(s string) bool {
	sl := len(s)
	rs := reverse(s)
	for i := 0; i < sl-1; i++ {
		subStr := s[i : i+2]
		if strings.Contains(s, subStr) && strings.Contains(rs, subStr) {
			return true
		}
	}
	return false
}

func reverse(s string) string {
	sl := len(s)
	sb := []byte(s)
	for i := 0; i < sl/2; i++ {
		sb[i], sb[sl-i-1] = sb[sl-i-1], sb[i]
	}
	return string(sb)
}

func countSubstrings(s string, c byte) int64 {
	sb := []byte(s)
	var count int64
	for i := 0; i < len(sb); i++ {
		if s[i] == c {
			count++
		}
	}
	return count*(count-1)/2+count
}

func minimumDeletions(word string, k int) int {
	wb := []byte(word)
	chMap := make(map[byte]int)
	for i := 0; i < len(wb); i++ {
		chMap[wb[i]]++ 
	}
	freqs := make([]int, 0, len(chMap))
	for _, value := range chMap {
		freqs = append(freqs, value)
	}
	
	sort.Ints(freqs)
	var ans int
	var minIndex int
	min , max := freqs[minIndex], freqs[len(freqs)-1]
	for max-min> k {
		if minIndex == len(freqs)-1 {
			break
		}
		diff := max-min
		maxSub := diff-k
		minSub := min
		if maxSub > minSub {
			ans+=maxSub
			break
		}
		ans+=minSub
		minIndex++
		min = freqs[minIndex]
	}
	return ans
}


  
