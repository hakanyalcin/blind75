package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
}

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if c == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' ||
				c == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' ||
				c == '}' && len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func isValidV2(s string) bool {
	stack := make([]rune, 0)
	bracketMap := map[rune]rune{
		')': '(', ']': '[', '}': '{',
	}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
