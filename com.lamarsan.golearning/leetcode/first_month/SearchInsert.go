package first_month

/**
 * description 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置
 *
 * @author lamar
 */
func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	} else if target > nums[len(nums)-1] {
		return len(nums)
	}
	for i, num := range nums {
		if num == target || target < num {
			return i
		}
	}
	return 0
}

func searchInsert2(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
