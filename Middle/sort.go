package main

// 一维数组排序
/*func sort(list []int) {
	//退出条件
	if len(list) <= 1 {
		return
	}
	i, j := 0, len(list)-1
	// 表示第一次比较的索引位置.
	index := 1
	// 第一次比较的参考值.这里选择第一个数
	key := list[0]
	if list[index] > key {
		list[i], list[j] = list[j], list[i]
		j-- //表示取大值跟末尾的数替换位置,使大于参考值的数在后面
	} else {
		list[i], list[index] = list[index], list[i]
		i++ //表示取小的值跟前面的替换位置,使小于参考值的数在前面
		index++
	}
	sort(list[:i])  //处理参考值前面的数组
	sort(list[i+1:]) //处理参考值后面的数组
}*/
