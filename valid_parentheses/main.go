package valid_parentheses

func isValid(s string) bool {
	stack := []byte{}
	if len(s)%2 > 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if s[i] == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else if s[i] == '}' && len(stack) > 0 && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
		} else if s[i] == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
