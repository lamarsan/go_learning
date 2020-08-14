package second_month

/**
 * description 给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，
 * 并且这些子字符串中的所有0和所有1都是组合在一起的。
 * 重复出现的子串要计算它们出现的次数。
 *
 * @author lamar
 * @date 2020/8/10 9:58 上午
 */
func countBinarySubstrings(s string) int {
	var ptr, last, ans int
	n := len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		ans += min(count, last)
		last = count
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
