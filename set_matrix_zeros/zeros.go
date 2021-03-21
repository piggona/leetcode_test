package main

func setZeros(matrix [][]int) {
	col0 := false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
