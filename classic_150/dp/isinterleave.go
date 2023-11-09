package dp

// 97. 交错字符串
// https://leetcode.cn/problems/interleaving-string/description/?envType=study-plan-v2&envId=top-interview-150

/*给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
注意：a + b 意味着字符串 a 和 b 连接。*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	s1l, s2l, s3l := len(s1), len(s2), len(s3)
	if s1l+s2l != s3l {
		return false
	}
	dp := make([][]bool, s1l+1)
	for i := range dp {
		dp[i] = make([]bool, s2l+1)
	}
	dp[0][0] = true
	for i := 0; i <= s1l; i++ {
		for j := 0; j <= s2l; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return dp[s1l][s2l]
}
