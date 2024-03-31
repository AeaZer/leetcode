package _0240331

import "strconv"

func sumOfTheDigitsOfHarshadNumber(x int) int {
	_, sum := intToSlice(x)
	if x%sum != 0 {
		return -1
	}
	return sum
}

func intToSlice(x int) ([]int, int) {
	xs := strconv.FormatInt(int64(x), 10)
	nums := make([]int, 0, len(xs))
	var sum int64
	for i := 0; i < len(xs); i++ {
		s := xs[i : i+1]
		pv, _ := strconv.ParseInt(s, 10, 64)
		sum += pv
		nums = append(nums, int(pv))
	}
	return nums, int(sum)
}

func maxBottlesDrunk(numBottles int, numExchange int) int {
	if numBottles < numExchange {
		return numBottles
	}
	ans := numBottles
	var emptyBottles int
	for emptyBottles+numBottles >= numExchange {
		emptyBottles++
		numBottles--
		if emptyBottles == numExchange {
			numExchange++
			numBottles++
			ans++
			emptyBottles = 0
		}
	}
	return ans
}

func countAlternatingSubarrays(nums []int) int64 {
	var ans int
	for i := 0; i < len(nums); {
		l := 1
		pre := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == pre {
				break
			}
			l++
			pre = nums[j]
		}
		i += l
		ans += (l + 1) * l / 2
	}
	return int64(ans)
}
