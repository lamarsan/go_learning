package first_month

/**
* description 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
* 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
*
* @author lamar
* @date 2020/7/22 9:43 上午
 */
func minArray(numbers []int) int {
	if numbers == nil || len(numbers) <= 0 {
		return 0
	}
	begin := numbers[0]
	for _, number := range numbers {
		if number < begin {
			return number
		}
	}
	return begin
}

func minArray2(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		pivot := low + (high-low)/2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}
