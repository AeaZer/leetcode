package _0240324

import "math"

func maximumLengthSubstring(s string) int {
	sbs := []byte(s)
	var ans int
	for i := 0; i < len(sbs); i++ {
		var numsRead [26]int
		j := i
		for ; j < len(sbs); j++ {
			index := sbs[j] - 'a'
			if numsRead[index] >= 2 {
				break
			}
			numsRead[index]++
		}
		if j-i > ans {
			ans = j - i
		}
	}
	return ans
}

func minOperations(k int) int {
	if k == 1 {
		return 0
	}
	minOperrate := math.MaxInt
	for i := 1; i < k; i++ {
		a := k / i
		sum := a - 1 + i-1
		b := k % i
		if b > 0 {
			sum++
		}
		if minOperrate > sum {
			minOperrate = sum
		}
	}
	return minOperrate
}

func mostFrequentIDs(nums []int, freq []int) []int64 {
	ans := make([]int64, len(nums))
	frequentIDMap := make(map[int]int64)
	maxKey, maxValue := -1, int64(-1)
	for i := 0; i < len(nums); i++ {
		currentNum := frequentIDMap[nums[i]] + int64(freq[i])
		frequentIDMap[nums[i]] = currentNum
		if nums[i] == maxKey && freq[i] >= 0 {
			ans[i] = currentNum
			maxValue = currentNum
			continue
		}
		if nums[i] != maxKey && currentNum <= maxValue {
			ans[i] = maxValue
			continue
		}
		maxKey, maxValue = maxInMap(frequentIDMap)
		ans[i] = maxValue
	}
	return ans
}

func maxInMap(m map[int]int64) (maxIntKey int, maxValue int64) {
	for key, value := range m {
		if maxValue < value {
			maxValue = value
			maxIntKey = key
		}
	}
	return
}