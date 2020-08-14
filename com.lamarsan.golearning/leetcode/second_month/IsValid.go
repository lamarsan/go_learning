package second_month

/**
 * description 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 * 有效字符串需满足：
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 【简单】
 *
 * @author lamar
 * @date 2020/8/14 9:44 上午
 */
type stack struct {
	value []rune
}

func (s *stack) get() rune {
	if len(s.value) == 0 {
		return '0'
	}
	n := len(s.value)
	result := s.value[n-1]
	s.value = s.value[:n-1]
	return result
}

func (s *stack) push(v rune) {
	s.value = append(s.value, v)
}

func isValid(s string) bool {
	r := []rune(s)
	a := 0
	b := 0
	c := 0
	last := stack{}
	for _, r2 := range r {
		if a < 0 || b < 0 || c < 0 {
			return false
		}
		if r2 == '(' {
			a = a + 1
			last.push('(')
		}
		if r2 == '[' {
			b = b + 1
			last.push('[')
		}
		if r2 == '{' {
			c = c + 1
			last.push('{')
		}
		if r2 == ')' {
			if last.get() != '(' {
				return false
			}
			a = a - 1
		}
		if r2 == ']' {
			if last.get() != '[' {
				return false
			}
			b = b - 1
		}
		if r2 == '}' {
			if last.get() != '{' {
				return false
			}
			c = c - 1
		}
	}
	if last.value != nil && len(last.value) != 0 {
		return false
	}
	return true
}
