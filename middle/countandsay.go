package main

import "fmt"

func countAndSay(n int) string {
	reverseFun(n, "1")
	return ""
}

func reverseFun(n int, str string) []string {
	res := make([]string, 0, len(str)*2)
	if n > 0 {
		for i := 0; i < len(str); i++ {
			strs := make([]string, 0, len(str))
			if i == 0 {
				strs = append(strs, str[i:i+1])
			} else {
				if str[i:i+1] == strs[0] {
					strs = append(strs, str[i:i+1])
				} else {
					res = append(res, []string{fmt.Sprint(len(strs)), strs[0]}...)
					i += len(strs)
				}
			}
		}
	}
	return res
}
