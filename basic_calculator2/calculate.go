package main

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = s + " "
	stack := []byte{}
	numStack := []int{}
	temp := 0
	var operation func(op byte)
	operation = func(op byte) {
		switch op {
		case '+':
			numStack[len(numStack)-2] = numStack[len(numStack)-1] + numStack[len(numStack)-2]
		case '-':
			numStack[len(numStack)-2] = numStack[len(numStack)-2] - numStack[len(numStack)-1]
		case '*':
			numStack[len(numStack)-2] = numStack[len(numStack)-2] * numStack[len(numStack)-1]
		case '/':
			numStack[len(numStack)-2] = numStack[len(numStack)-2] / numStack[len(numStack)-1]
		}
		numStack = numStack[:len(numStack)-1]
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num, _ := strconv.Atoi(string(s[i]))
			temp = temp*10 + num
		} else {
			if i-1 >= 0 && s[i-1] >= '0' && s[i-1] <= '9' {
				numStack = append(numStack, temp)
				temp = 0
			}
			if len(stack) != 0 {
				switch s[i] {
				case '+', '-':
					for len(stack) != 0 {
						op := stack[len(stack)-1]
						stack = stack[:len(stack)-1]
						operation(op)
					}
				case '*', '/':
					for len(stack) != 0 && (stack[len(stack)-1] == '*' || stack[len(stack)-1] == '/') {
						op := stack[len(stack)-1]
						stack = stack[:len(stack)-1]
						operation(op)
					}
				}
			}
			if s[i] != ' ' {
				stack = append(stack, s[i])
			}
		}
	}
	for len(stack) != 0 {
		operation(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return numStack[0]
}

func main() {
	fmt.Println(calculate("1+1+1"))
}
