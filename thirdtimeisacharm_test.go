package leetcode

func ThirdTimeIsACharm(str string) string {
	count := 0
	s := ""
	if len(str) < 3 {
		return "\n"
	}
	for i := 0; i < len(str); i++ {
		if count == 2 {
			s += string(str[i])
			count = 0
		} else {
			count++
		}

	}
	return s + "\n"
}
