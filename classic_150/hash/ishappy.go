package hash

/*编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。*/

// https://leetcode.cn/problems/happy-number/

func isHappy(n int) bool {
	m := make(map[int]bool)
	for {
		sum := pow2(n)
		if m[sum] {
			return false
		}
		if sum == 1 {
			break
		}
		n, m[sum] = sum, true
	}
	return true
}

func pow2(num int) int {
	var sum int
	for num != 0 {
		sum += (num % 10) * (num % 10)
		num /= 10
	}
	return sum
}
