package string_array

import "strings"

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
