package hash

/*给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
如果可以，返回 true ；否则返回 false 。
magazine 中的每个字符只能在 ransomNote 中使用一次。
*/

// https://leetcode.cn/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	magazineBytesMap := make(map[byte]uint, 0)
	for i := range magazine {
		magazineBytesMap[magazine[i]]++
	}
	for i := range ransomNote {
		// 不存在或者字符不够用
		if _, ok := magazineBytesMap[ransomNote[i]]; !ok || magazineBytesMap[ransomNote[i]] == 0 {
			return false
		}
		magazineBytesMap[ransomNote[i]]--
	}
	return true
}
