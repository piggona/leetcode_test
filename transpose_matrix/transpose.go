package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	w := len(matrix)
	if w == 0 {
		return matrix
	}
	l := len(matrix[0])

	result := make([][]int, l)
	for i := 0; i < l; i++ {
		result[i] = make([]int, w)
		for j := 0; j < w; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(transpose(matrix))
}
