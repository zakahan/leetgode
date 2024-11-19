package problem_79

// 这个没直接做出来，看了下题解
func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[0] {
			if dps(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dps(board [][]byte, word string, i int, j int, k int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true // 说明找到了
	}
	board[i][j] = ' '
	res := dps(board, word, i-1, j, k+1) || dps(board, word, i+1, j, k+1) || dps(board, word, i, j-1, k+1) || dps(board, word, i, j+1, k+1)
	// 回退
	board[i][j] = word[k]
	return res
}
