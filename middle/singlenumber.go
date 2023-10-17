package main

func singleNumber(nums []int) []int {
	sum := 0
	for _, num := range nums {
		sum ^= num
	}
	b := sum & -sum
	num1, num2 := 0, 0
	for _, num := range nums {
		if num&b > 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}
	return []int{num1, num2}
}
