package main

const incrementLoop = iota + 1

type matrixParams struct {
	matrixNum int
	finalXY   [2]int
	x         int
	y         int
	xLoop     *LoopParams
	yLoop     *LoopParams
}

type LoopParams struct {
	Increment int
	LoopNum   int
}

func newMatrixParams(matrixNum int) *matrixParams {
	m := new(matrixParams)
	m.matrixNum = matrixNum
	var finalXY [2]int
	if matrixNum%2 == 0 {
		finalXY[0] = matrixNum%2 + 1
		finalXY[1] = matrixNum%2 - 1
	} else {
		finalXY[0] = (matrixNum + 1) % 2
		finalXY[1] = (matrixNum + 1) % 2
	}
	m.finalXY = finalXY
	return m
}

func (p *matrixParams) generateLoop() [][]int {
	res := make([][]int, p.matrixNum)
	for i, _ := range res {
		subArray := make([]int, p.matrixNum)
		for j := 0; j < p.matrixNum; j++ {
			subArray[j] = 0
		}
		res[i] = subArray
	}
	count := 0
	for count <= p.matrixNum*p.matrixNum {

	}
	return res
}

func generateMatrix(n int) [][]int {
	return newMatrixParams(n).generateLoop()
}
