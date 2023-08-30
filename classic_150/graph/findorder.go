package graph

// 210. 课程表 II
// https://leetcode.cn/problems/course-schedule-ii/

// 现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites 
// ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

// 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
// 返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。

func findOrder(numCourses int, prerequisites [][]int) []int {
	courseMap := make(map[int][]int, 0)
	inDegrees := make([]int, numCourses)

	// 创建拓扑图
	for _, prerequisite := range prerequisites {
		course, nextCourse := prerequisite[0], prerequisite[1]
		courseMap[nextCourse] = append(courseMap[nextCourse], course)
		inDegrees[course]++
	}

	queue := make([]int, 0, numCourses)
	for course, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, course)
		}
	}

	courses := make([]int, 0, numCourses)
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		courses = append(courses, course)
		if len(courses) > numCourses {
			return []int{}
		}

		for _, neighbor := range courseMap[course] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// 没有环的情况会有入度不为 0，根据这个特性来判断是否有环
	for _, v := range inDegrees {
		if v != 0 {
			return []int{}
		}
	}

	return courses
}
