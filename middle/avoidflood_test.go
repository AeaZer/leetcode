package main

import "testing"

func TestAvoidFlood(t *testing.T) {
	flood := avoidFlood([]int{69, 0, 0, 0, 69})
	t.Log(flood)
}
