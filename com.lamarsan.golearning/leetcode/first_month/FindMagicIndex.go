package first_month

/**
 * description 魔术索引。 在数组A[0...n-1]中，有所谓的魔术索引，满足条件A[i] = i。
 * 给定一个有序整数数组，编写一种方法找出魔术索引，若有的话，在数组A中找出一个魔术索引，如果没有，则返回-1。
 * 若有多个魔术索引，返回索引值最小的一个。
 *
 * @author lamar
 * @date 2020/7/31 9:33 上午
 */
func findMagicIndex(nums []int) int {
	for i, num := range nums {
		if i == num {
			return i
		}
	}
	return -1
}

func findMagicIndex2(nums []int) int {
	return getAnswer(nums, 0, len(nums)-1)
}

func getAnswer(nums []int, left, right int) int {
	if left > right {
		return -1
	}
	mid := (right-left)/2 + left
	leftAnswer := getAnswer(nums, left, mid-1)
	if leftAnswer != -1 {
		return leftAnswer
	} else if nums[mid] == mid {
		return mid
	}
	return getAnswer(nums, mid+1, right)
}
