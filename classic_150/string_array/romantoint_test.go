package string_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	assert := assert.New(t)

	res := romanToInt("LVIII")
	assert.Equal(58, res)

	res = romanToInt("MCMXCIV")
	assert.Equal(1994, res)
}

func TestIntToRoman(t *testing.T) {
	assert := assert.New(t)

	roman := intToRoman(1994)
	assert.Equal("MCMXCIV", roman)

	roman = intToRoman(58)
	assert.Equal("LVIII", roman)
}
