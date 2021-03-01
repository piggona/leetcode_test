package main

import "fmt"

type NumArray struct {
	sumArray []int
}

func Constructor(nums []int) NumArray {
	sumArray := make([]int, len(nums))
	sum := 0
	for i := 0; i < len(nums); i++ {
		sumArray[i] = sum + nums[i]
		sum = sumArray[i]
	}
	return NumArray{
		sumArray: sumArray,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if len(this.sumArray) == 0 {
		return 0
	}
	if i == 0 {
		return this.sumArray[j]
	}
	return this.sumArray[j] - this.sumArray[i-1]
}

func main() {
	array := Constructor([]int{})
	fmt.Println(array.SumRange(0, 0))
}
