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

	aliceSizes := []int {1,2,5}
	bobSizes := []int {2,4}
	fmt.Println(fairCandySwap(aliceSizes,bobSizes))
}

/**
正统二分 704
 */
func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for {
		middle := left + (right - left) / 2
		fmt.Println("left:",left,"right:",right,"middle:",middle)
		// 左闭右闭 左边可以等于右边 所以左边大于右边的时候就可以退出
		if nums[middle] > target {
			right = middle - 1
		}
		if nums[middle] < target {
			left = middle + 1
		}
		fmt.Println("left:",left,"right:",right)
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
	fmt.Println("diff:",diff)
	// 即 b要交换的数大一点
	if diff < 0 {
		diff = diff * -1
		for _,v := range bobSizes {
			if isContain(aliceSizes, v - diff) {
				return [] int {v -diff, v}
			}
		}
	}

	for _,v := range aliceSizes {
		if isContain(bobSizes, v - diff) {
			return [] int {v, v -diff}
		}
	}

	return []int {}
}




func countArr(nums [] int) int{
	ans := 0
	for _,v  := range nums {
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





















