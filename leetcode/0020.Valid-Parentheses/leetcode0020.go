package leetcode

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, c := range s {
		if (c == '(') || (c == '[') || (c == '{') {
			stack = append(stack, c)
		} else if len(stack) != 0 && (((c == ')') && stack[len(stack)-1] == '(') || ((c == ']') && stack[len(stack)-1] == '[') || ((c == '}') && stack[len(stack)-1] == '{')) {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
