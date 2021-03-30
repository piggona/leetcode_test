package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] <= target && matrix[i][len(matrix[i])-1] >= target {
			return binarySearch(matrix[i], target)
		}
	}
	return false
}

func binarySearch(list []int, target int) bool {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		if list[middle] == target {
			return true
		}
		if list[middle] > target {
			right = middle - 1
		}
		if list[middle] < target {
			left = middle + 1
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 15))
}
