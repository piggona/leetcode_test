package main

import (
	"fmt"
	"sort"
	"time"
)

/*

 */

// O(n^2)
func twoSum_violence(nums []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}
	return result
}

// O(nlogn) + 2O(n) = O(nlog(n))
func twoSum_sort(nums []int, target int) []int {
	result := make([]int, 2)
	numMap := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		numMap[i] = i
	}
	sumStruct := &MapNum{
		Nums:   nums,
		NumMap: numMap,
	}
	sort.Sort(sumStruct)

	i := 0
	j := sumStruct.Len() - 1
	for i < j {
		if sumStruct.GetValue(i)+sumStruct.GetValue(j) < target {
			i++
		} else if sumStruct.GetValue(i)+sumStruct.GetValue(j) > target {
			j--
		} else {
			result[0] = sumStruct.NumMap[i]
			result[1] = sumStruct.NumMap[j]
			return result
		}
	}
	return []int{}
}

type MapNum struct {
	Nums   []int
	NumMap []int
}

func (m *MapNum) GetValue(i int) int {
	return m.Nums[m.NumMap[i]]
}

func (m *MapNum) Len() int {
	return len(m.Nums)
}

func (m *MapNum) Less(i, j int) bool {
	return m.GetValue(i) < m.GetValue(j)
}

func (m *MapNum) Swap(i, j int) {
	m.NumMap[i], m.NumMap[j] = m.NumMap[j], m.NumMap[i]
}

// O(n)
func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		pos, ok := numMap[target-nums[i]]
		if ok {
			result[0] = pos
			result[1] = i
			return result
		}
		numMap[nums[i]] = i
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 5}
	target := 9
	start := time.Now()
	result_violence := twoSum_violence(nums, target)
	fmt.Printf("result: %v, time: %d\n", result_violence, time.Since(start))
	start = time.Now()
	result_sort := twoSum_sort(nums, target)
	fmt.Printf("result: %v, time: %d\n", result_sort, time.Since(start))
	start = time.Now()
	result := twoSum(nums, target)
	fmt.Printf("result: %v,time: %d\n", result, time.Since(start))
}
