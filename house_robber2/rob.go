package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	if len(nums) < 3 {
		return max(nums...)
	}
	var gorob func(nums []int) int
	gorob = func(nums []int) int {
		first, second := nums[0], max(nums[0], nums[1])
		for _, v := range nums[2:] {
			first, second = second, max(first+v, second)
		}
		return second
	}
	return max(gorob(nums[1:]), gorob(nums[:len(nums)-1]))
}

func max(nums ...int) int {
	var max = math.MinInt32
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	fmt.Println(rob([]int{1, 3, 1, 3, 100}))
}
