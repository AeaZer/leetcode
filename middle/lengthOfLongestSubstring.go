package main

/*
// lengthOfLongestSubstring
func lengthOfLongestSubstring(s string) int {
	bytes := []byte(s)
	maxLength := 0
	byteLength := len(bytes)
	for i := 0; i < byteLength; i++ {
		count := 0
		j := 0
		if j < byteLength {
			if bytes[i] != bytes[j] {
				count++
				j++
				continue
			}
			if maxLength < count {
				maxLength = count
				break
			}
		}
	}

	return maxLength
}
*/
