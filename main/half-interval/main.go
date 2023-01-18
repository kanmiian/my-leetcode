package main

import (
	"fmt"
)

/**
二分法
*/
func main() {

	//nums := []int {2,5}
	//target := 5
	//fmt.Println(search(nums,target))

	// 公平的糖果
	//aliceSizes := []int {1,2,5}
	//bobSizes := []int {2,4}
	//fmt.Println(fairCandySwap(aliceSizes,bobSizes))

	// 268. 丢失的数字
	//nums := []int {9,6,4,2,3,5,7,0,1}
	//fmt.Println(missingNumber(nums))

	// 34
	nums := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(searchInsert(nums, target))
}

/**
1802. 有界数组中指定下标处的最大值
*/
func maxValue(n int, index int, maxSum int) int {
	left, right :=  1, maxSum

	for {
		if left >= right {
			break
		}
		mid := (left + right + 1) / 2
		if valid(mid,n,index,maxSum) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}

func mulCount(max int, len int) int {
	if len + 1 <max {
		small := max - len
		return (max -1 + small) * len / 2
	} else {
		ones := len - (max - 1)
		return max * (max - 1) / 2 +ones
	}
}

func valid(mid int, n int, index int, maxSum int) bool {
	left := index
	right := n - index - 1
	return mid+mulCount(mid, left)+mulCount(mid, right) <= maxSum
}

/**
正统二分 704
*/
func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for {
		middle := left + (right-left)/2
		fmt.Println("left:", left, "right:", right, "middle:", middle)
		// 左闭右闭 左边可以等于右边 所以左边大于右边的时候就可以退出
		if nums[middle] > target {
			right = middle - 1
		}
		if nums[middle] < target {
			left = middle + 1
		}
		fmt.Println("left:", left, "right:", right)
		if nums[middle] == target {
			return middle
		}
		if left > right {
			break
		}
	}

	return -1
}

/**
888 公平的糖果交换
*/
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	// 获取交换后的数组和值
	aliceAns := countArr(aliceSizes)
	bobAns := countArr(bobSizes)
	// 和差值 = 交换的数的差值
	diff := (aliceAns - bobAns) / 2
	fmt.Println("diff:", diff)
	// 即 b要交换的数大一点
	if diff < 0 {
		diff = diff * -1
		for _, v := range bobSizes {
			if isContain(aliceSizes, v-diff) {
				return []int{v - diff, v}
			}
		}
	}

	for _, v := range aliceSizes {
		if isContain(bobSizes, v-diff) {
			return []int{v, v - diff}
		}
	}

	return []int{}
}

func countArr(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans = ans + v
	}

	return ans
}

func isContain(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

/**
268. 丢失的数字
*/
func missingNumber(nums []int) int {
	length := len(nums)
	//sort.Ints(nums)
	//left := nums[0]
	//right := nums[length - 1]
	//if left != 0 {
	//	return 0
	//}
	//if right != length {
	//	return length
	//}
	//for {
	//	middle := left + (right - left) / 2
	//	actual:= nums[middle]
	//	fmt.Println("middle:",middle,"actual:",actual)
	//	if middle == actual {
	//		left = middle + 1
	//	}
	//	if middle < actual {
	//		right = middle - 1
	//	}
	//	fmt.Println("left:",left,"right:",right)
	//	if left > right {
	//		break
	//	}
	//	fmt.Println("left:",left,"right:",right)
	//}
	//return left
	//hasMap := map[int]bool{}
	//for _, v := range nums {
	//	hasMap[v] = true
	//}
	//for i := 0; ;i ++ {
	//	if !hasMap[i] {
	//		return i
	//	}
	//}
	// = =直接求和算差值最快也最优....
	total := length * (length + 1) / 2
	arrSum := 0
	for _, v := range nums {
		arrSum += v
	}

	return total - arrSum
}

/**
35. 搜索插入位置
*/
func searchInsert(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	for {
		if left > right {
			break
		}

		middle := left + (right-left)/2
		if nums[middle] > target {
			right = middle - 1
		}
		if nums[middle] < target {
			left = middle + 1
		}
		if nums[middle] == target {
			return middle
		}

	}
	return right + 1

	//for i,v := range nums {
	//	if v > target {
	//		return i
	//	}
	//}
	//
	//return len(nums)
}

/**
  34. 在排序数组中查找元素的第一个和最后一个位置
*/
func searchRange(nums []int, target int) []int {

	// 按照边界划分情况
	// 1. 在数组范围外（不存在) -1 -1
	// 2. 在数组范围内 (但不存在)  -1 -1
	// 3. 在数组范围内(存在) 返回位置

	// todo 为什么左右边界需要单独二分法查找
	// 左边是第一个等于target的位置 右边是最后一个大于target的位置 找不到就是-1

	return []int{-1, -1}
}

/**
寻找第一个等于target的值
*/
func getLeftBorder(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	leftBorder := -2
	for {
		middle := left + (right-left)/2
		if left > right {
			break
		}
		if nums[middle] > target {
			right = middle - 1
			leftBorder = right
		}
		if nums[middle] < target {
			left = middle + 1
		}
		// todo 怎么判断是否第一个出现的？
		if nums[middle] == target {
			right = middle - 1
			leftBorder = right
		}
	}

	return leftBorder
}

func getRightBorder(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	rightBorder := -1
	for {
		if left > right {
			break
		}
		middle := left + (right-left)/2
		if nums[middle] > target {

		}
	}

	return rightBorder
}
