package first_month

/**
 * description 给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
 * 函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
 *
 * @author lamar
 * @date 2020/7/20 9:07 上午
 */
func twoSum(numbers []int, target int) []int {
	var low, high int
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if target-numbers[i] == numbers[j] {
				low = i + 1
				high = j + 1
			}
		}
	}
	return []int{low, high}
}

func twoSum2(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low + 1, high + 1}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return []int{-1, -1}
}
