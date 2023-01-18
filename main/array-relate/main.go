package main

import (
	"fmt"
	"strings"
)

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
	//nums := []int{5, 7, 7, 8,  10}
	//target := 8
	//searchRange(nums, target)

	//operations := []string {"X++","++X","--X","X--"}
	//fmt.Println(finalValueAfterOperations(operations))

	//nums := []int{3, 2, 2, 3}
	//val := 3
	//fmt.Println(removeElement(nums, val))
	//twoOutOfThree();
}

// x的平方根
func mySqrt(x int) int {
	return 1
}

//  4. 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 整个为一个数组后排序取中间的一位

	//
	//middle1 := nums1

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
	//len := len(nums)
	//left, right := 0, len-1
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

/**
2011. 执行操作后的变量值
*/
func finalValueAfterOperations(operations []string) int {
	ans := 0
	for _, v := range operations {
		if strings.Contains(v, "+") {
			ans++
		}
		if strings.Contains(v, "-") {
			ans--
		}
	}

	return ans
}

func removeElement(nums []int, val int) int {
	size := len(nums)


	//for i := 0; i < size; i++ {
	//	if nums[i] == val {
	//		for j := i + 1; j < size; j++ {
	//			nums[j-1] = nums[j]
	//		}
	//		i--
	//		size--
	//	}
	//}





	return size
}

/**
	2032. 至少在两个数组中出现的值
 */
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	get := func(nums []int) (s [101]int) {
		for _, v := range nums {
			s[v] = 1
		}
		return
	}
	var ans []int
	s1, s2, s3 := get(nums1), get(nums2), get(nums3)
	for i := 1; i <= 100; i++ {
		if s1[i]+s2[i]+s3[i] > 1 {
			ans = append(ans, i)
		}
	}
	return ans
}



















