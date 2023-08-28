package graph

const (
	statusVisited = iota + 1
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	pl := len(prerequisites)
	if pl == 0 {
		return true
	}

	var isCircle bool
	courseMap := make(map[int]bool)
	visited := make([]int, pl)
	relyMap := make(map[int]bool)
	visited[0] = statusVisited
	var bfs func(cur []int)
	bfs = func(cur []int) {
		// 有环
		if isCircle {
			return
		}
		// 有依赖
		for i := 0; i < pl; i++ {
			if cur[1] == prerequisites[i][0] {
				if relyMap[prerequisites[i][1]] {
					isCircle = true
					return
				}
				// 剪枝
				if courseMap[prerequisites[i][1]] {
					return
				}
				visited[i] = statusVisited
				relyMap[prerequisites[i][1]] = true
				courseMap[prerequisites[i][0]] = true
				bfs(prerequisites[i])
			}
		}
		relyMap = make(map[int]bool)
		// 没有依赖
		for i := 0; i < pl; i++ {
			if visited[i] != statusVisited {
				relyMap[prerequisites[i][0]] = true
				relyMap[prerequisites[i][1]] = true
				courseMap[prerequisites[i][0]] = true
				courseMap[prerequisites[i][1]] = true
				visited[i] = statusVisited
				bfs(prerequisites[i])
			}
		}
	}
	relyMap[prerequisites[0][0]] = true
	relyMap[prerequisites[0][1]] = true
	courseMap[prerequisites[0][0]] = true
	courseMap[prerequisites[0][1]] = true
	bfs(prerequisites[0])

	return !isCircle && numCourses >= len(courseMap)
}
