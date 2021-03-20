package main

import "strconv"

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			stack[len(stack)-2] = num1 + num2
			stack = stack[:len(stack)-1]
		case "-":
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			stack[len(stack)-2] = num2 - num1
			stack = stack[:len(stack)-1]
		case "*":
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			stack[len(stack)-2] = num2 * num1
			stack = stack[:len(stack)-1]
		case "/":
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			stack[len(stack)-2] = num2 / num1
			stack = stack[:len(stack)-1]
		default:
			num, err := strconv.Atoi(tokens[i])
			if err != nil {
				panic(err)
			}
			stack = append(stack, num)
		}
	}
	return stack[0]
}
