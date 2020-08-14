package second_month

/**
 * description 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，
 * 它们的乘积也表示为字符串形式。
 *
 * @author lamar
 * @date 2020/8/13 8:40 下午
 */
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1 := []rune(reverseString(num1))
	n2 := []rune(reverseString(num2))
	finalResult := "0"
	for i, r := range n2 {
		number1 := int(r - '0')
		var result []rune
		flag := 0
		for _, r2 := range n1 {
			number2 := int(r2 - '0')
			tmp := number1 * number2
			result = append(result, rune(((flag+tmp)%10)+'0'))
			flag = (tmp + flag) / 10
		}
		if flag != 0 {
			result = append(result, rune(flag+'0'))
		}
		result = []rune(reverseString(string(result)))
		for k := 0; k < i; k++ {
			result = append(result, '0')
		}
		finalResult = addStrings(finalResult, string(result))
	}
	return finalResult
}

func addStrings(num1 string, num2 string) string {
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
