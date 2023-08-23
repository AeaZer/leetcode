package stack

import "strconv"

/*给你一个字符串数组 tokens ，表示一个根据 逆波兰表示法 表示的算术表达式。

请你计算该表达式。返回一个表示表达式值的整数。

注意：

有效的算符为 '+'、'-'、'*' 和 '/' 。
每个操作数（运算对象）都可以是一个整数或者另一个表达式。
两个整数之间的除法总是 向零截断 。
表达式中不含除零运算。
输入是一个根据逆波兰表示法表示的算术表达式。
答案及所有中间计算结果可以用 32 位 整数表示。*/

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := range tokens {
		token, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, token)
			continue
		}
		stackTop, stackTopNext := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		switch tokens[i] {
		case "+":
			stack = append(stack, stackTop+stackTopNext)
		case "-":
			stack = append(stack, stackTop-stackTopNext)
		case "*":
			stack = append(stack, stackTop*stackTopNext)
		default: // case "/":
			stack = append(stack, stackTop/stackTopNext)
		}
	}
	return stack[0]
}
