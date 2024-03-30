package _0240316

import (
	"sort"
	"strconv"
)

func sumOfEncryptedInt(nums []int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		ans += encrypt(nums[i])
	}
	return ans
}

func encrypt(x int) int {
	xs := strconv.FormatInt(int64(x), 10)
	maxBit := '0'
	for _, value := range xs {
		if value > maxBit {
			maxBit = value
		}
	}
	var ansStr string 
	for i := 0; i < len(xs); i++ {
		ansStr += string(maxBit)
	}
	ans , _ := strconv.Atoi(ansStr) 
	return ans
}

type mark struct {
	index int
	value int
	marked bool
}

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	marks := make([]mark, 0, len(nums))
	var sum int
	for i := 0; i < len(nums); i++ {
		marks = append(marks, mark{
			index: i,
			value: nums[i],
		})
		sum += nums[i]
	}
	sort.Slice(marks, func(i, j int) bool {
		if marks[i].value == marks[j].value {
			return marks[i].index < marks[j].index
		}
		return marks[i].value < marks[j].value
	})
	// map[index]*mark
	seachMap := make(map[int]*mark, len(marks))
	for i := 0; i < len(marks); i++ {
		seachMap[marks[i].index] = &marks[i]
	}
	ans := make([]int64, 0)
	var f int
	for _, value := range queries {
		singleMarkIndex, multiMarkSum := value[0],value[1]
		singleMark := seachMap[singleMarkIndex]
		if !singleMark.marked {
			singleMark.marked = true
			sum-=singleMark.value
		}
		for multiMarkSum!=0 {
			if f >= len(marks) {
				break
			}
			if marks[f].marked {
				f++
				continue
			} 
			marks[f].marked = true
			sum-=marks[f].value
			f++
			multiMarkSum--
		}
		ans = append(ans, int64(sum))
	}
	return ans
}