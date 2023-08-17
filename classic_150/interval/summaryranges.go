package interval

import (
	"fmt"
	"strconv"
)

/*给定一个  无重复元素 的 有序 整数数组 nums 。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，
nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b*/

// https://leetcode.cn/problems/summary-ranges/

// 快慢双指针，一个指针指向开始的地方另外一个指向结束的地方
func summaryRanges(nums []int) []string {
	numsL := len(nums)
	if numsL == 0 {
		return []string{}
	}
	res := make([]string, 0)
	slow, fast, step := 0, 0, 0
	for {
		fast++
		step++
		// 防止数组越界，也是唯一可以结束 for 循环的逻辑块
		if fast >= numsL {
			if step == 1 {
				res = append(res, strconv.Itoa(nums[slow]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[slow], nums[fast-1]))
			}
			break
		}
		// 如果是连续的，那就继续好了
		if nums[fast]-nums[slow] == step {
			continue
		}
		// 中间部分，如果只走了一步，说明当前元素没有下一个节点与它连续
		if step == 1 {
			res = append(res, strconv.Itoa(nums[slow]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[slow], nums[fast-1]))
		}
		// 重置双指针，进入下一个遍历
		slow = fast
		// 重置步数
		step = 0
	}
	return res
}
