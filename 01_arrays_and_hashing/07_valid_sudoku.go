package arraysandhashing

func isValidSudoku(board [][]byte) bool {
	rows := [9]uint16{}
	cols := [9]uint16{}
	fields := [9]uint16{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			val := board[i][j] - '0'

			if rows[i]&(1<<val) > 0 {
				return false
			}
			rows[i] |= (1 << val)

			if cols[j]&(1<<val) > 0 {
				return false
			}
			cols[j] |= (1 << val)

			fieldIdx := i/3*3 + j/3
			if fields[fieldIdx]&(1<<val) > 0 {
				return false
			}
			fields[fieldIdx] |= (1 << val)
		}
	}

	return true
}
