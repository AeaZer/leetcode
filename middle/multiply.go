package main

import (
	"fmt"
	"math"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	var res int64
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num1Int, _ := strconv.ParseInt(num1, 10, 64)
	length := len(num2)
	for index := range num2 {
		pareInt, _ := strconv.ParseInt(num2[index:index+1], 10, 64)
		if pareInt == 0 {
			continue
		}

		res += int64(math.Pow10(length-index-1)) * pareInt * num1Int
	}

	return fmt.Sprint(res)
}
