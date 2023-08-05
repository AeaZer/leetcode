package string_array

/*在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
*/

// https://leetcode.cn/problems/gas-station/

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 1 && gas[0] == cost[0] {
		return 0
	}

	for i := 0; i < len(gas); i++ {
		if gas[i] <= cost[i] {
			continue
		}
		if runCircle(gas, cost, i) {
			return i
		}
	}
	return -1
}

func runCircle(gas []int, cost []int, startIndex int) bool {
	var (
		length      = len(gas)
		stationThey int
		gasHave     = gas[startIndex] // 从 startIndex 出发，有开始 station 的 gas
	)
	for stationThey != length {
		if gasHave < cost[startIndex%length] {
			return false
		}
		gasHave += gas[(startIndex+1)%length] - cost[startIndex%length]
		startIndex++
		stationThey++
	}
	return true
}
