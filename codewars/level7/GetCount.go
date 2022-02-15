package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetCount(str string) (count int) {
	arr := "aeiou"
	// Enter solution here
	str = strings.TrimSpace(str)
	fmt.Println(str)
	for _, element := range str {

		if strings.Contains(arr, strconv.Itoa(int(element))) {
			count++
		}
	}

	return count
}
