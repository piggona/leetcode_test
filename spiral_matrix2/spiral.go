package main

import "fmt"

func generateMatrix(num int) [][]int {
	result := make([][]int, num)
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, num)
	}
	count := 1
	direction := 0
	h := 0
	w := 0
	for h >= 0 && h < num && w >= 0 && w < num && result[h][w] == 0 {
		result[h][w] = count
		count++
		switch direction {
		case 0:
			if w+1 > num-1 || result[h][w+1] != 0 {
				direction = (direction + 1) & 3
				h++
			} else {
				w++
			}
		case 1:
			if h+1 > num-1 || result[h+1][w] != 0 {
				direction = (direction + 1) & 3
				w--
			} else {
				h++
			}
		case 2:
			if w-1 < 0 || result[h][w-1] != 0 {
				direction = (direction + 1) & 3
				h--
			} else {
				w--
			}
		case 3:
			if h-1 < 0 || result[h-1][w] != 0 {
				direction = (direction + 1) & 3
				w++
			} else {
				h--
			}
		}
	}
	return result
}

type pair struct{ x, y int }

var dirs = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	row, col, dirIdx := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := dirs[dirIdx]
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIdx = (dirIdx + 1) % 4 // 顺时针旋转至下一个方向
			dir = dirs[dirIdx]
		}
		row += dir.x
		col += dir.y
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
}
