package bitcup

// 67. 二进制求和
// 给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和
// https://leetcode.cn/problems/add-binary/description/?envType=study-plan-v2&envId=top-interview-150

const (
	byte0 = '0'
)

func addBinary(a string, b string) string {
	al, bl := len(a), len(b)
	var builder []byte
	if al > bl {
		builder = make([]byte, al+1)
	} else {
		builder = make([]byte, bl+1)
	}
	aBytes, bBytes := []byte(a), []byte(b)
	rest := byte('0')
	i, j := al-1, bl-1
	bIndex := len(builder) - 1
	for i >= 0 && j >= 0 {
		now := rest + aBytes[i] + bBytes[j] - 3*byte0
		rest = now/2 + byte0
		builder[bIndex] = (now % 2) + byte0
		bIndex--
		i--
		j--
	}
	for i >= 0 {
		now := rest + aBytes[i] - 2*byte0
		rest = now/2 + byte0
		builder[bIndex] = (now % 2) + byte0
		bIndex--
		i--
	}
	for j >= 0 {
		now := rest + bBytes[j] - 2*byte0
		rest = now/2 + byte0
		builder[bIndex] = (now % 2) + byte0
		bIndex--
		j--
	}
	builder[0] = rest
	if rest == byte0 {
		builder = builder[1:]
	}
	return string(builder)
}
