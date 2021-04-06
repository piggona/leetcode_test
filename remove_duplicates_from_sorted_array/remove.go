package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	var left, right int
	var pos int
	left = 0
	right = 2
	pos = 2
	for right < len(nums) {
		if nums[left] != nums[right] {
			nums[pos] = nums[right]
			pos++
			left++
			right++
		} else {
			right++
		}
	}
	return pos
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	length := removeDuplicates(nums)
	fmt.Println(nums)
	fmt.Println(length)
}
