package main

import "fmt"

/*
1 <= words.length <= 10^5
4 <= words[i].length <= 50
1 <= puzzles.length <= 10^4
puzzles[i].length == 7
words[i][j], puzzles[i][j] are English lowercase letters.
Each puzzles[i] doesn't contain repeated characters.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle
*/
func findNumOfValidWords(words []string, puzzles []string) []int {
	result := make([]int, len(puzzles))
	puzzleMap := make(map[int]int)
	for i := 0; i < len(words); i++ {
		puzzleMap[wordBit(words[i])]++
	}
	for i := 0; i < len(puzzles); i++ {
		puzzleBit := wordBit(puzzles[i])
		first := wordBit(string(puzzles[i][0]))

		n := puzzleBit
		for n > 0 {
			if (n&first > 0) && (puzzleMap[n] > 0) {
				result[i] += puzzleMap[n]
			}
			n = (n - 1) & puzzleBit
		}
	}
	return result
}

func wordBit(word string) int {
	var result int
	for i := 0; i < len(word); i++ {
		offset := word[i] - 'a'
		temp := 1 << int(offset)
		result = result | temp
	}
	return result
}

func main() {
	words := []string{
		"aaaa",
		"asas",
		"able",
		"ability",
		"actt",
		"actor",
		"access",
	}
	puzzles := []string{
		"aboveyz",
		"abrodyz",
		"abslute",
		"absoryz",
		"actresz",
		"gaswxyz",
	}
	fmt.Println(findNumOfValidWords(words, puzzles))
}
