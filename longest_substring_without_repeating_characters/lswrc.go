package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// 一个窗口，head,tail
	// 一个map[string]int，记录每个字母的位置+1
	// tail向后滑动，当碰到有字母重复的，head向前移动，直到碰到与其相同的元素
	if len(s) == 0 {
		return 0
	}
	var (
		left     = 0
		right    = 0
		result   = 0
		countmap = make(map[byte]int)
	)
	for right < len(s) {
		if countmap[s[right]] == 0 {
			countmap[s[right]]++
			right++
			if right-left > result {
				result = right - left
			}
		} else {
			countmap[s[left]]--
			left++
		}
	}
	return result
}

// 当穷举的元素有限，使用hashmap提升效率
func lengthOfLongestSubstring_bitmap(s string) int {
	freq := make([]int, 128)
	var res = 0
	start, end := 0, -1
	for start < len(s) {
		if end+1 < len(s) && freq[s[end+1]] == 0 {
			end++
			freq[s[end]]++
		} else {
			freq[s[start]]--
			start++
		}
		res = max(res, end-start+1)
	}
	return res
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	s := "abba"
	fmt.Println(lengthOfLongestSubstring(s))
}
