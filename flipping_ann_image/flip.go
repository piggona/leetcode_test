package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		l := 0
		r := len(A[i]) - 1
		for l < r {
			A[i][l] = abs(A[i][l] - 1)
			A[i][r] = abs(A[i][r] - 1)
			A[i][l], A[i][r] = A[i][r], A[i][l]
			l++
			r--
		}
		if l == r {
			A[i][l] = abs(A[i][l] - 1)
		}
	}
	return A
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	A := [][]int{
		{1},
	}
	fmt.Println(flipAndInvertImage(A))
}
