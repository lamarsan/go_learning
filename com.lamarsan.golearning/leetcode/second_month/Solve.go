package second_month

/**
 * description 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
 * 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
 *
 * @author lamar
 * @date 2020/8/11 9:47 上午
 */
func solve(board [][]byte) {
	for i, bytes := range board {
		for j, b := range bytes {
			flag := make([][]bool, len(board))
			for i := 0; i < len(board); i++ {
				flag[i] = make([]bool, len(board[i]))
			}
			if b == 'O' && canBeReplaced(board, flag, i, j) {
				board[i][j] = 'X'
			}
		}
	}
}

func canBeReplaced(board [][]byte, flag [][]bool, i int, j int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
		return true
	}
	if board[i][j] == 'X' || flag[i][j] == true {
		return true
	}
	if i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1 {
		if board[i][j] == 'O' {
			return false
		}
	}
	flag[i][j] = true
	return canBeReplaced(board, flag, i-1, j) && canBeReplaced(board, flag, i+1, j) &&
		canBeReplaced(board, flag, i, j-1) && canBeReplaced(board, flag, i, j+1)
}

// ---------------------------------------------------------------------------------------
var n, m int

func solve2(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m = len(board), len(board[0])
	for i := 0; i < n; i++ {
		dfs2(board, i, 0)
		dfs2(board, i, m-1)
	}
	for i := 1; i < m-1; i++ {
		dfs2(board, 0, i)
		dfs2(board, n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs2(board [][]byte, x, y int) {
	if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs2(board, x+1, y)
	dfs2(board, x-1, y)
	dfs2(board, x, y+1)
	dfs2(board, x, y-1)
}
