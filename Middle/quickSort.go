package main

func quickSort(nums *[]int, l, r int) {
	left, right := l, r
	if left <= right {
		temp := (*nums)[left]
		for left != right {
			for right > left {
				right--
				(*nums)[left] = (*nums)[right]
			}
			for left < right {
				left++
				(*nums)[right] = (*nums)[left]
			}
		}
		(*nums)[right] = temp
		quickSort(nums, l, left-1)
		quickSort(nums, right+1, r)
	}
}
