package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	type args struct {
		matrix Matrix
	}
	tests := []struct {
		name     string
		args     args
		expected Matrix
	}{
		{
			"1",
			args{
				Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			Matrix{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate(tt.args.matrix)
		})
		if !reflect.DeepEqual(tt.args.matrix, tt.expected) {
			t.Error("Not equal!", tt.args.matrix, tt.expected)
		}
	}
}

func BenchmarkRotate(b *testing.B) {
	const matrixLen = 100
	for i := 0; i < b.N; i++ {
		matrix := make([][]int, matrixLen)
		for r := 0; r < matrixLen; r++ {
			matrix[r] = make([]int, matrixLen)
			for c := 0; c < matrixLen; c++ {
				matrix[r][c] = rand.Intn(10000)
			}
		}
		Rotate(matrix)
	}
}
