package main

import (
	"fmt"
	"reflect"
)

func main() {

	image := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			fmt.Printf("%v -> %v\n",
				coordinateType{i, j},
				image.calcNextCoord(coordinateType{i, j}))
		}
	}

	fmt.Println(image)
	Rotate(image)
	fmt.Println(image)
	fmt.Println(reflect.DeepEqual(image, Matrix{
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
	fmt.Println(image)
	Rotate(image)
	fmt.Println(image)
	fmt.Println(reflect.DeepEqual(image, Matrix{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}))

	image = [][]int{
		{1},
	}
	Rotate(image)
	fmt.Println(reflect.DeepEqual(image, Matrix{
		{1},
	}))

	image = [][]int{{1, 2}, {3, 4}}
	Rotate(image)
	fmt.Println(reflect.DeepEqual(image, Matrix{{3, 1}, {4, 2}}))

}
