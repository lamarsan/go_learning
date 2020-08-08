package first_month

/**
 * description 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
 * 你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
 * 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
 *
 * @author lamar
 * @date 2020/7/27 9:34 上午
 */
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) || len(t) == 0 {
		return false
	}
	r := []rune(s)
	j := 0
	for _, i := range t {
		if j == len(s) {
			break
		}
		if r[j] == i {
			j++
		}
	}
	return j == len(s)
}
