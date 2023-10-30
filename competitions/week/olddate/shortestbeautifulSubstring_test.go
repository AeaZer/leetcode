package olddate

import "testing"

func TestShortestBeautifulSubstring(t *testing.T) {
	substring := shortestBeautifulSubstring("011", 2)
	t.Log(substring)
}
