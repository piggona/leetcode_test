package main

import "fmt"

type NumMatrix struct {
	sumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{}
	}
	sum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		sum[i] = make([]int, len(matrix[0]))
	}
	sum[0][0] = matrix[0][0]
	for i := 1; i < len(matrix); i++ {
		sum[i][0] = sum[i-1][0] + matrix[i][0]
	}
	for j := 1; j < len(matrix[0]); j++ {
		sum[0][j] = sum[0][j-1] + matrix[0][j]
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			sum[i][j] = matrix[i][j] + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
		}
	}
	return NumMatrix{
		sumMatrix: sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.sumMatrix[row2][col2]
	}
	if row1 == 0 {
		return this.sumMatrix[row2][col2] - this.sumMatrix[row2][col1-1]
	}
	if col1 == 0 {
		return this.sumMatrix[row2][col2] - this.sumMatrix[row1-1][col2]
	}
	return this.sumMatrix[row2][col2] - this.sumMatrix[row1-1][col2] - this.sumMatrix[row2][col1-1] + this.sumMatrix[row1-1][col1-1]
}

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	numMatrix := Constructor(matrix)
	fmt.Println(numMatrix.SumRegion(1, 2, 2, 4))
}
