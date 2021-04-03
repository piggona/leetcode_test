package main

import "fmt"

func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

func max(nums ...int) int {
	result := 0
	for _, value := range nums {
		if result < value {
			result = value
		}
	}
	return result
}

func min(nums ...int) int {
	result := 0
	for _, value := range nums {
		if result > value {
			result = value
		}
	}
	return result
}

func main() {
	fmt.Println(longestCommonSubsequence("bl", "yby"))
}
