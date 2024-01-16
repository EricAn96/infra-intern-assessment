package main
import "fmt"

func SolveSudoku(board [][]int) [][]int {
	// Make a quick lookup map for rows
	rows := make([]map[int]bool, 9)

	// Make a quick lookup map for columns
	columns := make([]map[int]bool, 9)

	// Make a quick lookup map for each cells(or boxes).
	// Each cell represents 3x3 sub-board which the basic sudoku rule is applied.
	// ASCII example:
	// . . . | . . . | . . .
	// . 0 . | . 1 . | . 2 .
	// . . . | . . . | . . .
	// ------+-------+------
	// . . . | . . . | . . .
	// . 3 . | . 4 . | . 5 .
	// . . . | . . . | . . .
	// ------+-------+------
	// . . . | . . . | . . .
	// . 6 . | . 7 . | . 8 .
	// . . . | . . . | . . .
	cells := make([]map[int]bool, 9)

	// Initialize each map in rows, columns, and cells.
	// The boolean value being true represents that number being present in that map.
	// For example, if rows[8][1] = true, that means the 9th row of this sudoku board has 1 somewhere.
	for i := 0; i < 9; i++ {
		rows[i] = make(map[int]bool)
		columns[i] = make(map[int]bool)
		cells[i] = make(map[int]bool)
	}

	// Since some of the blocks in the given sudoku board already have values, mark them as true, so that the algorithm is aware of them.
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != 0 {
				columns[c][board[r][c]] = true
				rows[r][board[r][c]] = true
				cells[(r/3)*3+(c/3)][board[r][c]] = true
			}
		}
	}

	// Main algorithm where this sudoku board is solved.
	backtrack(board, rows, columns, cells, 0, 0)
	return board
}

// This function uses a recursive backtracking algorithm that iterates from top left block to bottom right block.
// Each row will be 'solved' or 'attempted' before moving on to the next row.
func backtrack(board [][]int, rows, columns, cells []map[int]bool, r, c int) bool {
	// if row == 9, that means the algorithm has reached the end
	if r == 9 {
		return true
	}

	// if the board at position (r, c) already has a value, go to the next block.
	if board[r][c] != 0 {
		if c < 8 {
			return backtrack(board, rows, columns, cells, r, c+1)
		} else {
			return backtrack(board, rows, columns, cells, r+1, 0)
		}
	}

	cellNum := (r/3)*3 + (c/3)

	// The function tries all numbers between 1-9 inclusive
	for num := 1; num <= 9; num++ {
		// This if case checks if marking this block with the num is allowed
		if rows[r][num] || columns[c][num] || cells[cellNum][num] {
			continue
		}

		// Mark this block with the num
		rows[r][num] = true
		columns[c][num] = true
		cells[cellNum][num] = true
		board[r][c] = num

		// If there was a case where ALL the subsequent calls succeeded, this call will return true.
		var nextCall bool
		if c < 8 {
			nextCall = backtrack(board, rows, columns, cells, r, c+1)
		} else {
			nextCall = backtrack(board, rows, columns, cells, r+1, 0)
		}
		if nextCall {
			return true
		}

		// If there weren't any case where ALL the subsequent calls succeeded, revoke the markings done, and try the next number.
		rows[r][num] = false
		columns[c][num] = false
		cells[cellNum][num] = false
		board[r][c] = 0
	}
	return false
}

func printBoard(grid [][]int) {
	for i, row := range grid {
		if i%3 == 0 && i != 0 {
			fmt.Println("------+-------+------")
		}
		for j, num := range row {
			if j%3 == 0 && j != 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}

func main() {
	sudokuBoard := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	solved := SolveSudoku(sudokuBoard)
	printBoard(solved)
}