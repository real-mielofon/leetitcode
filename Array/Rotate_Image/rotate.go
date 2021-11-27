package main

import "fmt"

type coordinateType struct {
	r int
	c int
}

func (c coordinateType) String() string {
	return fmt.Sprintf("r: %d, c: %d", c.r, c.c)
}

type Matrix [][]int

func (m Matrix) String() string {
	s := ""
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			s = s + fmt.Sprintf("%d, ", m[i][j])
		}
		s = s + fmt.Sprintf("\n")
	}
	return s
}

func (m Matrix) calcNextCoord(coord coordinateType) coordinateType {
	r := coord.c
	c := len(m) - 1 - coord.r
	return coordinateType{r, c}
}

func Rotate(matrix Matrix) {
	rotatedCells := make(map[coordinateType]bool)

	rotateCell := func(coord coordinateType) {
		nextCoord := matrix.calcNextCoord(coord)
		curCellValue := matrix[coord.r][coord.c]
		for {
			if _, ok := rotatedCells[nextCoord]; ok {
				break
			}
			saveCell := matrix[nextCoord.r][nextCoord.c]
			matrix[nextCoord.r][nextCoord.c] = curCellValue
			rotatedCells[nextCoord] = true
			nextCoord = matrix.calcNextCoord(nextCoord)
			curCellValue = saveCell
		}
	}

	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix)-1; j++ {
			if _, ok := rotatedCells[coordinateType{i, j}]; !ok {
				rotateCell(coordinateType{i, j})
			}
		}
	}
}
