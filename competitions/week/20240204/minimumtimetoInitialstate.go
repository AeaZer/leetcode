package _0240204

/*给你一个下标从 0 开始的字符串 word 和一个整数 k 。

在每一秒，你必须执行以下操作：

移除 word 的前 k 个字符。
在 word 的末尾添加 k 个任意字符。
注意 添加的字符不必和移除的字符相同。但是，必须在每一秒钟都执行 两种 操作。

返回将 word 恢复到其 初始 状态所需的 最短 时间（该时间必须大于零）。*/

// abacaba 3
func minimumTimeToInitialState(word string, k int) int {
	wl := len(word)
	i := k
	var ans int
	for ; i < wl; i += k {
		ans++
		l := wl - i
		if word[:l] == word[i:] {
			break
		}
	}
	return ans
}
