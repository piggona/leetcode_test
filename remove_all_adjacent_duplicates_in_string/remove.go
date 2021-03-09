package main

func removeDuplicates(s string) string {
	if len(s) < 2 {
		return s
	}
	i := 0
	j := 1
	for j < len(s) {
		if s[i] == s[j] {
			for i-1 >= 0 && j+1 < len(s) && s[i-1] == s[j+1] {
				i--
				j++
			}
			return removeDuplicates(s[:i] + s[j+1:])
		}
		i++
		j++
	}
	return s
}

func removeDuplicatesStack(s string) string {
	stack := []byte{}
	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
