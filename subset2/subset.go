package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
outer:
	for mask := 0; mask < 1<<n; mask++ {
		t := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					continue outer
				}
				t = append(t, v)
			}
		}
		ans = append(ans, append([]int(nil), t...))
	}
	return
}

func main() {
	fmt.Println(4 >> 2 & 1)
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
