package main

import (
	"fmt"
	"math"
)

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		op := false
		for len(stack) != 0 && stack[len(stack)-1] > nums[i] {
			op = true
			stack = stack[:len(stack)-1]
		}
		if op && len(stack) != 0 {
			return true
		}
		stack = append(stack, nums[i])
	}
	return false
}

func find132pattern2(nums []int) bool {
	n := len(nums)
	candidateK := []int{nums[n-1]}
	maxK := math.MinInt64

	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(candidateK) > 0 && nums[i] > candidateK[len(candidateK)-1] {
			maxK = candidateK[len(candidateK)-1]
			candidateK = candidateK[:len(candidateK)-1]
		}
		if nums[i] > maxK {
			candidateK = append(candidateK, nums[i])
		}
	}

	return false
}

func main() {
	fmt.Println(find132pattern([]int{3, 5, 0, 3, 4}))
	return
}
