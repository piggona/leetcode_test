package main

import "sort"

func shipWithinDays(weights []int, D int) int {
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	return left + sort.Search(right-left, func(x int) bool {
		x += left
		day := 1
		sum := 0
		for _, weight := range weights {
			if sum+weight > x {
				day++
				sum = 0
			}
			sum += weight
		}
		return day <= D
	})
}
