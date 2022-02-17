package main

func lengthOfLastWord(s string) (ans int) {
	size := len(s) - 1
	lastIndex := size
	for s[size] == ' ' {
		lastIndex--
	}
	for lastIndex >= 0 && s[lastIndex] != ' ' {
		ans++
		lastIndex--
	}
	return
}
