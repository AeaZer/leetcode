package matrix

/*根据 百度百科 ， 生命游戏 ，简称为 生命 ，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。

给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态： 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。给你 m x n 网格面板 board 的当前状态，返回下一个状态。
*/

// https://leetcode.cn/problems/game-of-life/

func gameOfLife(board [][]int) {
	n, m := len(board), len(board[0])
	tmp := make([][]int, n+2)
	for i := range tmp {
		tmp[i] = make([]int, m+2)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp[i+1][j+1] = board[i][j]
		}
	}

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			aliveCount := tmp[i][j] + tmp[i][j+1] + tmp[i][j+2] +
				tmp[i+1][j] + tmp[i+1][j+2] +
				tmp[i+2][j] + tmp[i+2][j+1] + tmp[i+2][j+2]
			if board[i][j] == 1 {
				switch aliveCount {
				case 2, 3:
					board[i][j] = 1
				default:
					board[i][j] = 0
				}
				continue
			}
			if board[i][j] == 0 && aliveCount == 3 {
				board[i][j] = 1
			}
		}
	}
}
