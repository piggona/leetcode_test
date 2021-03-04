package main

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Sort(twi(envelopes))
	tmp := make([]int, len(envelopes))
	var result int
	tmp[0] = 1
	for i := 1; i < len(envelopes); i++ {
		j := i - 1
		tmp[i] = 1
		for j >= 0 {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				if tmp[j]+1 > tmp[i] {
					tmp[i] = tmp[j] + 1
				}
			}
			j--
		}
		if tmp[i] > result {
			result = tmp[i]
		}
	}
	return result
}

type twi [][]int

func (t twi) Len() int {
	return len(t)
}

func (t twi) Less(i, j int) bool {
	num1 := (t)[i][0]
	num2 := (t)[j][0]
	if num1 == num2 {
		num1 = (t)[i][1]
		num2 = (t)[j][1]
	}
	if num1 < num2 {
		return true
	}
	return false
}

func (t twi) Swap(i, j int) {
	(t)[i], (t)[j] = (t)[j], (t)[i]
}

func main() {
	input := [][]int{
		{2, 4},
		{5, 2},
		{5, 5},
		{6, 3},
		{6, 7},
		{7, 4},
		{8, 5},
	}
	fmt.Println(maxEnvelopes(input))
}
