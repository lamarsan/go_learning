package second_month

/**
 * description 给定一组唯一的单词， 找出所有不同 的索引对(i, j)，使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。
 *
 * @author lamar
 * @date 2020/8/6 9:47 上午
 */
func palindromePairs(words []string) [][]int {
	if len(words) == 0 {
		return nil
	}
	var result [][]int
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			if isPalindrome(words[i], words[j]) {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func isPalindrome(a string, b string) bool {
	c := a + b
	n := len(c)
	for i := 0; i < n/2; i++ {
		if c[i] != c[n-i-1] {
			return false
		}
	}
	return true
}

//-----------------------------------------------------------------------------
// 上面超时了，所以换一种思路，如果s1+s2为回文串，len(s1)==len(s2)，那么s1为s2的回文串；
// 若len(s1)>len(s2)，那么可将s1拆为t1和t2，t1为s2的反转，t2本身为回文串；
// 若len(s1)<len(s2)，那么可将s2拆为t1和t2，t2为s1的反转，t1本身为回文串。
func palindromePairs2(words []string) [][]int {
	var wordsRev []string
	indices := map[string]int{}

	n := len(words)
	for _, word := range words {
		wordsRev = append(wordsRev, reverse(word))
	}
	for i := 0; i < n; i++ {
		indices[wordsRev[i]] = i
	}

	var ret [][]int
	for i := 0; i < n; i++ {
		word := words[i]
		m := len(word)
		if m == 0 {
			continue
		}
		for j := 0; j <= m; j++ {
			if isPalindrome2(word, j, m-1) {
				leftId := findWord(word, 0, j-1, indices)
				if leftId != -1 && leftId != i {
					ret = append(ret, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome2(word, 0, j-1) {
				rightId := findWord(word, j, m-1, indices)
				if rightId != -1 && rightId != i {
					ret = append(ret, []int{rightId, i})
				}
			}
		}
	}
	return ret
}

func findWord(s string, left, right int, indices map[string]int) int {
	if v, ok := indices[s[left:right+1]]; ok {
		return v
	}
	return -1
}

func isPalindrome2(s string, left, right int) bool {
	for i := 0; i < (right-left+1)/2; i++ {
		if s[left+i] != s[right-i] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	n := len(s)
	b := []byte(s)
	for i := 0; i < n/2; i++ {
		b[i], b[n-i-1] = b[n-i-1], b[i]
	}
	return string(b)
}
