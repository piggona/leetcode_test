package main

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	pos := 0
	temp := 0
	op := 1
	var cal func(s string) int
	cal = func(s string) int {
		result := 0
		for pos < len(s) {
			if s[pos] >= '0' && s[pos] <= '9' {
				number, _ := strconv.Atoi(string(s[pos]))
				temp = temp*10 + number
			} else if s[pos] != ' ' {
				result = result + op*temp
				temp = 0
				if s[pos] == '(' {
					sop := op
					op = 1
					pos++
					result = result + sop*cal(s)
					continue
				}
				op = 1
				if s[pos] == '-' {
					op = -1
				}
				if s[pos] == ')' {
					pos++
					return result
				}

			}
			pos++
		}
		return result + op*temp
	}
	return cal(s)
}

func main() {
	fmt.Println(calculate("99"))
}
