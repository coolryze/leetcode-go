package removealladjacentduplicatesinstring

func removeDuplicates(s string) string {
	stack := []rune{}

	for _, ch := range s {
		if len(stack) == 0 || stack[len(stack)-1] != ch {
			stack = append(stack, ch)
			continue
		}

		stack = stack[:len(stack)-1]
	}

	return string(stack)
}
