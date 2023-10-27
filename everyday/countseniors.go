package everyday

import "fmt"

func countSeniors(details []string) int {
	var ans int
	for i := range details {
		fmt.Println(string(details[i][11]))
		if details[i][11] > '6' || (details[i][11] >= '6' && details[i][12] > '0') {
			ans++
		}
	}
	return ans
}
