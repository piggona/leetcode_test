package main

import (
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}
	if len(preorder) == 0 {
		return false
	}
	preorderList := strings.Split(preorder, ",")
	count := make([]int, len(preorderList))
	var countFinder func(i int) bool
	countFinder = func(i int) bool {
		for i >= 0 {
			if count[i] > 0 {
				count[i]--
				if i != 0 && count[i] == 0 {
					return countFinder(i)
				}
				break
			}
			i--
		}
		if i < 0 {
			return false
		}
		return true
	}
	for i := 0; i < len(count); i++ {
		count[i] = 2
	}
	for pos := 0; pos < len(preorderList); pos++ {
		if preorderList[pos] == "#" {
			count[pos] = 0
			if !countFinder(pos) {
				return false
			}
		}
	}
	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isValidSerialization("#"))
}
