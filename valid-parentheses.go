package leetcode

func isValid(s string) bool {

	count := 0
	if len(s)%2 != 0 || s == "" {
		return false
	}
	mp := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}

	for i := 0; i < len(s); i++ {

		if s[i] == ')' || s[i] == '}' || s[i] == ']' {

			if len(stack) > 0 {
				if mp[rune(s[i])] == stack[len(stack)-1] {
					count--
					stack = stack[:len(stack)-1]
				} else {

					return false
				}
			} else {
				return false

			}

		} else {

			count++
			stack = append(stack, rune(s[i]))
		}

	}
	if count != 0 {
		return false
	}
	return true
}
