package string_array

import "strings"

/*罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1 。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。*/

// https://leetcode.cn/problems/roman-to-integer/

func romanToInt(s string) int {
	var intMap = map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var res int
	for i := 0; i < len(s); i++ {
		currentValue := intMap[s[i]]
		if i+1 < len(s) && currentValue < intMap[s[i+1]] {
			res -= currentValue
			continue
		}
		res += currentValue
	}
	return res
}

/*罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，
所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给你一个整数，将其转为罗马数字。*/

// https://leetcode.cn/problems/integer-to-roman/

func intToRoman(num int) string {
	var intMap = map[int]byte{
		1:    'I',
		5:    'V',
		10:   'X',
		50:   'L',
		100:  'C',
		500:  'D',
		1000: 'M',
	}
	var enumInt = []int{1000, 500, 100, 50, 10, 5, 1}
	multiTimes := make([]int, 7)
	for i := range enumInt {
		multiTimes[i] = num / enumInt[i]
		num -= multiTimes[i] * enumInt[i]
	}

	var builder strings.Builder
	for i := 0; i < 7; i++ {
		// 特殊情况
		if (i == 2 || i == 4 || i == 6) && multiTimes[i] == 4 {
			builder.WriteByte(intMap[enumInt[i]])
			builder.WriteByte(intMap[enumInt[i-1-multiTimes[i-1]]])
			continue
		}
		// 特殊情况不处理，等下一个节点处理
		if (i == 1 || i == 3 || i == 5) && multiTimes[i+1] == 4 {
			continue
		}
		for j := 0; j < multiTimes[i]; j++ {
			builder.WriteByte(intMap[enumInt[i]])
		}
	}
	return builder.String()
}
