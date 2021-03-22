package main

import "fmt"

func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		num = num & (num - 1)
		result++
	}
	return result
}

func main() {
	fmt.Println(hammingWeight(5))
}
