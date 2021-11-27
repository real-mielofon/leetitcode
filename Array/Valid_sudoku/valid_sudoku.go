package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	const (
		boardSize = 9
		cellSize  = 3
	)
	line := func(n int) bool {
		numbersInRange := make(map[byte]bool)
		for i := 0; i < boardSize; i++ {
			cell := board[i][n]
			if cell == '.' {
				continue
			}
			if _, ok := numbersInRange[cell]; ok {
				return false
			}
			numbersInRange[cell] = true
		}
		return true
	}
	row := func(n int) bool {
		numbersInRange := make(map[byte]bool)
		for i := 0; i < boardSize; i++ {
			cell := board[n][i]
			if cell == '.' {
				continue
			}
			if _, ok := numbersInRange[cell]; ok {
				return false
			}
			numbersInRange[cell] = true
		}
		return true
	}
	cells9 := func(n int) bool {
		numbersInRange := make(map[byte]bool)
		line := (n / cellSize) * cellSize
		row := (n % cellSize) * cellSize
		for i := 0; i < cellSize; i++ {
			for j := 0; j < cellSize; j++ {
				cell := board[line+i][row+j]
				if cell == '.' {
					continue
				}
				if _, ok := numbersInRange[cell]; ok {
					return false
				}
				numbersInRange[cell] = true
			}
		}
		return true
	}
	for i := 0; i < boardSize; i++ {
		result := line(i) && row(i) && cells9(i)
		if !result {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},

		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},

		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
	fmt.Println(isValidSudoku([][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},

		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},

		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
}
