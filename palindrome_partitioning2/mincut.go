package main

import "fmt"

func minCut(s string) int {
	partition := make([]int, len(s))
	if len(s) < 2 {
		return 0
	}
	pstr := make([][]bool, len(s))
	for i := 0; i < len(pstr); i++ {
		pstr[i] = make([]bool, len(s))
		for j := 0; j < len(s); j++ {
			pstr[i][j] = true
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			pstr[i][j] = s[i] == s[j] && pstr[i+1][j-1]
		}
	}
	for right := 1; right < len(s); right++ {
		temp := partition[right-1] + 1
		for left := right - 1; left > 0; left-- {
			if pstr[left][right] {
				temp = min(partition[left-1]+1, temp)
			}
		}
		if pstr[0][right] {
			temp = 0
		}
		partition[right] = temp
	}
	return partition[len(s)-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Println(minCut("aab"))
}
