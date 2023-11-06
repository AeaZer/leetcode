package _0231105

// 100115. 找到冠军 I
// https://leetcode.cn/problems/find-champion-i/description/
/*一场比赛中共有 n 支队伍，按从 0 到  n - 1 编号。

给你一个下标从 0 开始、大小为 n * n 的二维布尔矩阵 grid 。
对于满足 0 <= i, j <= n - 1 且 i != j 的所有 i, j ：如果 grid[i][j] == 1，那么 i 队比 j 队 强 ；否则，j 队比 i 队 强 。
在这场比赛中，如果不存在某支强于 a 队的队伍，则认为 a 队将会是 冠军 。
返回这场比赛中将会成为冠军的队伍。*/

/*输入：grid = [[0,1],[0,0]]
输出：0
解释：比赛中有两支队伍。
grid[0][1] == 1 表示 0 队比 1 队强。所以 0 队是冠军。*/

func findChampionI(grid [][]int) int {
	teamLength := len(grid)
	for i := 0; i < teamLength; i++ {
		var ans int
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				ans++
			}
		}
		if ans == teamLength-1 {
			return i
		}
	}
	return -1
}

// 100116. 找到冠军 II
// https://leetcode.cn/problems/find-champion-ii/description/
/*一场比赛中共有 n 支队伍，按从 0 到  n - 1 编号。每支队伍也是 有向无环图（DAG） 上的一个节点。
给你一个整数 n 和一个下标从 0 开始、长度为 m 的二维整数数组 edges 表示这个有向无环图
，其中 edges[i] = [ui, vi] 表示图中存在一条从 ui 队到 vi 队的有向边。
从 a 队到 b 队的有向边意味着 a 队比 b 队 强 ，也就是 b 队比 a 队 弱 。
在这场比赛中，如果不存在某支强于 a 队的队伍，则认为 a 队将会是 冠军 。
如果这场比赛存在 唯一 一个冠军，则返回将会成为冠军的队伍。否则，返回 -1 。*/

/*输入：n = 3, edges = [[0,1],[1,2]]
输出：0
解释：1 队比 0 队弱。2 队比 1 队弱。所以冠军是 0 队。*/

func findChampionII(n int, edges [][]int) int {
	m := make(map[int][]int, n)
	for _, edge := range edges {
		m[edge[1]] = append(m[edge[1]], edge[0])
	}
	var check int
	visited := make([]bool, n)
	for key := range m {
		if !visited[key] {
			check++
			visited[key] = true
		}
	}
	if n-check > 1 {
		return -1
	}
	var ans int
	for _, edge := range edges {
		if len(m[edge[0]]) == 0 {
			ans = edge[0]
		}
	}
	return ans
}
