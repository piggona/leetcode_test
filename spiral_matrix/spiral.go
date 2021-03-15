package main

import (
	"fmt"
	"math"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	result := make([]int, len(matrix)*len(matrix[0]))
	var move func(matrix [][]int)
	move = func(matrix [][]int) {
		direction := 0
		h := 0
		w := 0
		pos := 0
		for h >= 0 && w >= 0 && h < len(matrix) && w < len(matrix[0]) && matrix[h][w] != math.MinInt32 {
			result[pos] = matrix[h][w]
			matrix[h][w] = math.MinInt32
			pos++
			switch direction {
			case 0:
				if w+1 >= len(matrix[0]) || matrix[h][w+1] == math.MinInt32 {
					direction = (direction + 1) & 3
					h++
				} else {
					w++
				}
			case 1:
				if h+1 >= len(matrix) || matrix[h+1][w] == math.MinInt32 {
					direction = (direction + 1) & 3
					w--
				} else {
					h++
				}
			case 2:
				if w-1 < 0 || matrix[h][w-1] == math.MinInt32 {
					direction = (direction + 1) & 3
					h--
				} else {
					w--
				}
			case 3:
				if h-1 < 0 || matrix[h-1][w] == math.MinInt32 {
					direction = (direction + 1) & 3
					w++
				} else {
					h--
				}
			}
		}
	}
	move(matrix)
	return result
}

func main() {

	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println(spiralOrder(matrix))
	return
}
