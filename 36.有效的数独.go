/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	row := make([][]bool, 9)
	col := make([][]bool, 9)
	matrix := make([][]bool, 9)
	initArrays(row)
	initArrays(col)
	initArrays(matrix)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v == '.' {
				continue
			}
			v = v - '1'
			if row[i][v] {
				return false
			}
			row[i][v] = true

			if col[j][v] {
				return false
			}
			col[j][v] = true
			matrixNum := (i / 3) * 3 + j / 3
			if matrix[matrixNum][v] {
				return false
			}
			matrix[matrixNum][v] = true
		}
	}
	return true
}

func initArrays(arr [][]bool) {
	for i := 0; i < 9; i++ {
		arr[i] = make([]bool, 9)
	}
}
// @lc code=end

