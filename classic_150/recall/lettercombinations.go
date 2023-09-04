package recall

// 17. 电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/

// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

func letterCombinations(digits string) []string {
	var telMap = map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	digitsBytes := []byte(digits)
	bl := len(digitsBytes)
	res := make([]string, 0)
	var iterator func(key int)
	iterator = func(index int) {
		if index >= bl {
			return
		}
		values := telMap[digitsBytes[index]]
		if len(res) != 0 {
			cur := make([]string, 0, len(values)*len(res))
			for _, re := range res {
				for _, value := range values {
					cur = append(cur, re+value)
				}
			}
			res = cur
		} else {
			res = append(res, values...)
		}
		iterator(index + 1)
	}
	iterator(0)
	return res
}
