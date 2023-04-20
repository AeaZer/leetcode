package main

type combineTimeParams struct {
	startTime int64
	endTime   int64
}

func initData() []combineTimeParams {
	return []combineTimeParams{
		{
			startTime: 10,
			endTime:   100,
		},
		{
			startTime: 20,
			endTime:   50,
		},
		{
			startTime: 30,
			endTime:   80,
		},
		{
			startTime: 40,
			endTime:   70,
		},
		{
			startTime: 60,
			endTime:   90,
		},
	}
}

func maxCombine() int {
	timeParams := initData()
	mc := 0
	for i := 0; i < len(timeParams); i++ {
		tempCombineNum := 0
		for j := i + 1; j < len(timeParams); j++ {
			if timeParams[j-1].endTime > timeParams[j].startTime {
				tempCombineNum++
			}
		}
		if tempCombineNum > mc {
			mc = tempCombineNum
		}
	}
	return mc
}
