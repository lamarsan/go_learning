package first_month

import "strconv"

/**
 * description 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
 *
 * @author lamar
 * @date 2020/8/3 10:10 上午
 */
func addStrings(num1 string, num2 string) string {
	a, err := strconv.Atoi(num1)
	if err != nil {
		return ""
	}
	b, err2 := strconv.Atoi(num2)
	if err2 != nil {
		return ""
	}
	return strconv.Itoa(a + b)
}

func addStrings2(num1 string, num2 string) string {
	r1 := []rune(reverseString(num1))
	r2 := []rune(reverseString(num2))

	n := max(len(num1), len(num2))
	c := 0
	result := make([]rune, n+1)
	for i := 0; i < n; i++ {
		a := 0
		b := 0
		if len(r1) > i {
			a = int(r1[i] - '0')
		}
		if len(r2) > i {
			b = int(r2[i] - '0')
		}
		result[i] = rune((a+b+c)%10 + '0')
		if a+b+c >= 10 {
			c = 1
		} else {
			c = 0
		}
	}
	if c != 0 {
		result[n] = '1'
	} else {
		result = result[:n]
	}
	return reverseString(string(result))
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
