package main

// 2512. 奖励最顶尖的 K 名学生
// https://leetcode.cn/problems/reward-top-k-students/description/

/*给你两个字符串数组 positive_feedback 和 negative_feedback ，分别包含表示正面的和负面的词汇。不会 有单词同时是正面的和负面的。

一开始，每位学生分数为 0 。每个正面的单词会给学生的分数 加 3 分，每个负面的词会给学生的分数 减  1 分。

给你 n 个学生的评语，用一个下标从 0 开始的字符串数组 report 和一个下标从 0 开始的整数数组 student_id 表示，其中 student_id[i]
表示这名学生的 ID ，这名学生的评语是 report[i] 。每名学生的 ID 互不相同。

给你一个整数 k ，请你返回按照得分 从高到低 最顶尖的 k 名学生。如果有多名学生分数相同，ID 越小排名越前。*/

import (
	"sort"
	"strings"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	feedbackMap := make(map[string]int, len(positive_feedback)+len(negative_feedback))
	for _, s := range positive_feedback {
		feedbackMap[s] = 3
	}
	for _, s := range negative_feedback {
		feedbackMap[s] = -1
	}
	type pairs struct {
		id    int
		score int
	}
	students := make([]pairs, len(student_id))
	for i, s := range report {
		feedbacks := strings.Split(s, " ")
		var score int
		for _, fd := range feedbacks {
			score += feedbackMap[fd]
		}
		students[i] = pairs{
			id:    student_id[i],
			score: score,
		}
	}
	sort.Slice(students, func(i, j int) bool {
		if students[i].score != students[j].score {
			return students[i].score > students[j].score
		}
		return students[i].id < students[j].id
	})
	ids := make([]int, k)
	for i := 0; i < k; i++ {
		ids[i] = students[i].id
	}
	return ids
}
