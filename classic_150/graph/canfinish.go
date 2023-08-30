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
// 207. 课程表
// https://leetcode.cn/problems/course-schedule/

/* 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，
其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。 */

// canFinish2 拓扑排序
func canFinish2(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int) // 有向图 int => []int
	inDegrees := make([]int, numCourses) // 入度数组

	// 构建图和入度数组
	for _, prerequisite := range prerequisites {
		course, dep := prerequisite[0], prerequisite[1]
		graph[dep] = append(graph[dep], course)
		inDegrees[course]++
	}

	queue := make([]int, 0) //入度为 0 的数组
	for course, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, course)
		}
	}

	count := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		count++

		for _, neighbor := range graph[course] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return count == numCourses
}
