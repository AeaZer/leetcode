package bitcup

import "testing"

func TestAddBinary(t *testing.T) {
	binary := addBinary("111", "1")
	t.Log(binary)
}
