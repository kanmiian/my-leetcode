package main

import "fmt"

func main() {
	// 704 二分查找
	//nums := []int{-1,0,3,5,9,12}
	//target := 2
	//search(nums,target)

	// no 35
	//nums := []int{1, 3, 5, 6}
	//target := 7
	//searchInsert(nums, target)

	// 复写零
	//arr := []int{0, 1, 2, 3, 3, 0, 4, 2}
	//duplicateZeros(arr)
	//testArr()

	// 34 获取数字出现的最开始与末尾位置
	nums := []int{5, 7, 7, 8,  10}
	target := 8
	searchRange(nums, target)

}

// x的平方根
func mySqrt(x int) int {
	return 1
}


//  4. 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 整个为一个数组后排序取中间的一位

	//
	middle1 := nums1



	fmt.Println(float64(1), "result")
	return float64(1)
}


//
func searchRange(nums []int, target int) []int {
	//  获取左右边界 左右可以相同 代表只有一个
	rightBorder := getRightBorder(nums, target)
	leftBorder := getLeftBorder(nums, target)
	fmt.Println(rightBorder, "right", leftBorder, "left")
	// 判断三种情况 如果只出现一次，
	if nums[rightBorder-1] != target {
		return []int{-1, -1}
	}
	fmt.Println(leftBorder-1, rightBorder-1, "test")
	return []int{leftBorder + 1, rightBorder - 1}
}

/**
获取右边界
*/
func getRightBorder(nums []int, target int) int {
	len := len(nums)
	left, right := 0, len-1
	border := -1

	// Left = right => 左右边界相同的，全包
	for left < right {
		middle := left + ((right - left) / 2)
		if nums[middle] > target {
			right = middle
			border = right
		} else {
			left = middle
		}
	}

	return border
}

func getLeftBorder(nums []int, target int) int {
	len := len(nums)
	left, right := 0, len-1
	border := -1

	return border
}

func search(nums []int, target int) int {
	len := len(nums)
	left, right := 0, len-1
	result := -1
	middle := left + ((right - left) >> 1)
	for i := 0; i < len; i++ {
		if left <= right {
			middle = left + ((right - left) >> 1)
			if nums[middle] == target {
				result = middle
			}
			if nums[middle] > target {
				right = middle - 1
			}
			if nums[middle] < target {
				left = middle + 1
			}
		}
	}

	fmt.Println(result)
	return result
}

func searchInsert(nums []int, target int) int {
	len := len(nums)
	left, right := 0, len-1
	middle := left + ((right - left) >> 1)
	for i := 0; i < len; i++ {
		if left <= right {
			middle = left + ((right - left) >> 1)
			fmt.Println(left, right, middle)
			if nums[middle] == target {
				return middle
			}
			if nums[middle] < target {
				left = middle + 1
			}
			if nums[middle] > target {
				right = middle - 1
			}
		}
	}

	// 不存在于数组内
	fmt.Println(right+1, "11111", middle+1)
	return right + 1
}

func testArr() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	target := 2
	result := -1
	len := len(arr)

	// 左右闭合  left <= right
	left, right := 0, len-1
	middle := left + ((right - left) / 2)
	//for i := 0; i < len; i++ {
	//	if left > right {
	//		break
	//	}
	//	if left <= right {
	//		middle = left + ((right - left) / 2)
	//		fmt.Println(middle, left, right)
	//		if arr[middle] == target {
	//			result = middle
	//			break
	//		}
	//		if arr[middle] < target {
	//			left = middle + 1
	//		}
	//		if arr[middle] > target {
	//			right = middle - 1
	//		}
	//	}
	//}

	//  [) 的时候 left  < right
	for i := 0; i < len; i++ {
		if left < right {
			if arr[middle] == target {
				result = middle
				break
			}
			if arr[middle] < target {
				left = middle + 1
			}
			if arr[middle] > target {
				right = middle
			}

		} else {
			break
		}
	}
	fmt.Println(result)
}
