package main

import "fmt"

func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		if i&1 == 0 {
			result[i] = result[i/2]
		} else {
			result[i] = result[i-1] + 1
		}
	}
	return result
}

func main() {
	num := 5
	fmt.Println(countBits(num))
}
