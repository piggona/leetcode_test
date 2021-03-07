package main

import "fmt"

func partition(s string) [][]string {
	result := [][]string{}
	replication := make(map[string]bool)
	if len(s) == 0 {
		return [][]string{}
	}
	if isPalindrome(s) {
		result = append(result, []string{s})
	}
	for i := 1; i < len(s); i++ {
		res := merge(partition(s[:i]), partition(s[i:]))
		for _, temp := range res {
			mark := fmt.Sprintln(temp)
			if replication[mark] == false {
				result = append(result, temp)
				replication[mark] = true
			}
		}
	}
	return result
}

func isPalindrome(s string) bool {
	if len(s)%2 == 0 {
		i := len(s)/2 - 1
		j := len(s) / 2
		for i >= 0 {
			if s[i] != s[j] {
				return false
			}
			i--
			j++
		}
		return true
	}
	i := len(s)/2 - 1
	j := len(s)/2 + 1
	for i >= 0 {
		if s[i] != s[j] {
			return false
		}
		i--
		j++
	}
	return true
}

func merge(s1, s2 [][]string) [][]string {
	if len(s1) == 0 || len(s2) == 0 {
		return append(s1, s2...)
	}
	result := [][]string{}
	replication := make(map[string]bool)
	for i := 0; i < len(s2); i++ {
		for j := 0; j < len(s1); j++ {
			temp := append(s1[j], s2[i]...)
			mark := fmt.Sprintln(temp)
			if replication[mark] == false {
				result = append(result, temp)
				replication[mark] = true
			}
		}
	}
	return result
}

func main() {
	result := partition("aabcb")
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
