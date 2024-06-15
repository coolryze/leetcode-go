package validparentheses

func isValid(s string) bool {
	stack := []rune{}
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 || m[ch] != stack[len(stack)-1] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
