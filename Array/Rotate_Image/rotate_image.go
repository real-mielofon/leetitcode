package main

import (
	"fmt"
	"reflect"
)

func rotate(matrix [][]int) {
	l := len(matrix)

	type coordinateType struct {
		r int
		c int
	}
	rotatedCells := make(map[coordinateType]bool)

	calcNextCoord := func(coord coordinateType) coordinateType {
		return coordinateType{l - coord.r, coord.c}
	}

	rotateCell := func(coord coordinateType) {
		nextCoord := calcNextCoord(coord)
		curCellValue := matrix[coord.r][coord.c]
		for {
			if _, ok := rotatedCells[nextCoord]; ok {
				break
			}
			saveCell := matrix[nextCoord.r][nextCoord.c]
			matrix[nextCoord.r][nextCoord.c] = curCellValue
			rotatedCells[nextCoord] = true
			nextCoord = calcNextCoord(nextCoord)
			curCellValue = saveCell
		}
	}

	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1; j++ {
			if _, ok := rotatedCells[coordinateType{i, j}]; !ok {
				rotateCell(coordinateType{i, j})
			}
		}
	}
}

func main() {

	image := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(image)
	fmt.Println(reflect.DeepEqual(image, [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}))

	image = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(image)
	fmt.Println(reflect.DeepEqual(image, [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}))

	image = [][]int{
		{1},
	}
	rotate(image)
	fmt.Println(reflect.DeepEqual(image, [][]int{
		{1},
	}))

	image = [][]int{{1, 2}, {3, 4}}
	rotate(image)
	fmt.Println(reflect.DeepEqual(image, [][]int{{3, 1}, {4, 2}}))

}
