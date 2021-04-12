package main

import (
	"fmt"
	"sort"
	"strconv"
)

type SortNums []int

func (s SortNums) Len() int {
	return len(s)
}

func (s SortNums) Less(i, j int) bool {
	var mult func(num int) int
	mult = func(num int) int {
		if num == 0 {
			return 10
		}
		base := 1
		for num != 0 {
			num /= 10
			base *= 10
		}
		return base
	}
	return (s[i]*mult(s[j]) + s[j]) > (s[j]*mult(s[i]) + s[i])
}

func (s SortNums) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func largestNumber(nums []int) string {
	sort.Sort(SortNums(nums))
	result := ""
	for _, num := range nums {
		result += strconv.Itoa(num)
	}
	if result[0] == '0' {
		return "0"
	}
	return result
}

func main() {
	fmt.Println(largestNumber([]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println(strconv.Atoi("00"))
}
