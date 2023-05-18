package main

import (
	"fmt"
	"math"
)

var hexBase10Map = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"A": 10,
	"B": 11,
	"C": 12,
	"D": 13,
	"E": 14,
	"F": 15,
}

const hexMulti float64 = 16

func main() {
	fmt.Println(hexToInt("0xAA"))
}

func hexToInt(hex string) int32 {
	var intRes int32
	var y float64
	for i := len(hex) - 1; i > 1; i-- {
		intRes += int32(hexBase10Map[string(hex[i])] * int(math.Pow(hexMulti, y)))
		y++
	}
	return intRes
}
