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
		if v, ok := symbolMap[s[i]]; ok {
			currentL := len(stack)
			if currentL == 0 || stack[currentL-1] != v {
				return false
			}
			stack = stack[:currentL-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}
