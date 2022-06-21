package main

import "fmt"

func main()  {
	// 7014 二分查找
	nums := []int{-1,0,3,5,9,12}
	target := 2
	search(nums,target)


	// no 35
	//nums := []int{1, 3, 5, 6}
	//target := 7
	//searchInsert(nums, target)

	// 复写零
	//arr := []int{0, 1, 2, 3, 3, 0, 4, 2}
	//duplicateZeros(arr)
	//testArr()

	// 34 获取数字出现的最开始与末尾位置
	//nums := []int{1, 3, 5, 6}
	//target := 7
	//searchRange(nums, target)
}

// todo
func searchRange(nums []int, target int) []int {
	var result []int
	len := len(nums)
	left, right := 0, len-1
	middle := left + ((right - left) >> 1)
	// 顺序排列
	for i := 0; i < len; i ++ {
		if left <= right {
			middle = left + ((right - left) >> 1)
			if nums[middle] == target {
				result = append(result, middle)
			}
			if nums[middle] > target {
				right = middle - 1
			}
			if nums[middle] < target {
				left = middle + 1
			}
		}
	}

	return result
}


func search(nums []int, target int) int {
	len := len(nums)
	left, right := 0, len - 1
	result := -1
	middle := left + ((right - left) >> 1)
	for i := 0; i < len; i ++ {
		if left <= right {
			middle = left + ((right - left) >> 1)
			if nums[middle] == target {
				result =  middle
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
			fmt.Println(left,right,middle)
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
	fmt.Println(right + 1,"11111", middle + 1)
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