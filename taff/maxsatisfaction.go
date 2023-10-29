package taff

import "sort"

// 1402. 做菜顺序
// https://leetcode.cn/problems/reducing-dishes/description/?envType=daily-question&envId=2023-10-22

/*一个厨师收集了他 n 道菜的满意程度 satisfaction ，这个厨师做出每道菜的时间都是 1 单位时间。

一道菜的 「 like-time 系数 」定义为烹饪这道菜结束的时间（包含之前每道菜所花费的时间）乘以这道菜的满意程度，也就是 time[i]*satisfaction[i] 。

返回厨师在准备了一定数量的菜肴后可以获得的最大 like-time 系数 总和。

你可以按 任意 顺序安排做菜的顺序，你也可以选择放弃做某些菜来获得更大的总和。*/

func maxSatisfaction(satisfaction []int) int {
	negative, positive := make([]int, 0), make([]int, 0)
	var positiveNumsSum int
	for i := range satisfaction {
		score := satisfaction[i]
		if score < 0 {
			negative = append(negative, score)
			continue
		}
		positiveNumsSum += score
		positive = append(positive, score)
	}
	sort.Slice(negative, func(i, j int) bool {
		return negative[i] > negative[j]
	})
	sort.Ints(positive)

	negativeIndex := 0
	for ; negativeIndex < len(negative); negativeIndex++ {
		if negative[negativeIndex]*(negativeIndex+1)+positiveNumsSum < 0 {
			break
		}
	}

	var ans int
	for i := 0; i < negativeIndex; i++ {
		ans += negative[i] * (i + 1)
	}
	for i := 0; i < len(positive); i++ {
		ans += (negativeIndex + i + 1) * positive[i]
	}

	return ans
}
