package stack

/*给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

1.左括号必须用相同类型的右括号闭合。
2.左括号必须以正确的顺序闭合。
3.每个右括号都有一个对应的相同类型的左括号。*/

// 20. 有效的括号
// https://leetcode.cn/problems/valid-parentheses/

/*1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成*/

func isValid(s string) bool {
	sL := len(s)
	if sL%2 != 0 {
		return false
	}

	symbolMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for i := range s {
		// 当元素是右括号时判断栈中的第一个元素是否是对应的左括号
		if v, ok := symbolMap[s[i]]; ok {
			currentL := len(stack)
			// 栈为空或者第一个元素不是右括号，说明不是有效数组
			if currentL == 0 || stack[currentL-1] != v {
				return false
			}
			// 检查没问题就 cut 掉，相当于出栈
			stack = stack[:currentL-1]
			continue
		}
		// 推入左括号入栈数组
		stack = append(stack, s[i])
	}
	// 判断栈中的左括号是否消耗完成 == 左右括号成对存在
	return len(stack) == 0
}
