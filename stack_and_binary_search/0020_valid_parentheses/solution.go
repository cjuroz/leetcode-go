package validparentheses

func IsValidMap(s string) bool {
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]byte, 0, len(s))

	for i := range s {
		ch := s[i]
		if openBracket, isClosing := pairs[ch]; isClosing {
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}

func IsValid(s string) bool {
	stack := make([]byte, 0, len(s))

	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (top == '(' && s[i] == ')') ||
				(top == '{' && s[i] == '}') ||
				(top == '[' && s[i] == ']') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
