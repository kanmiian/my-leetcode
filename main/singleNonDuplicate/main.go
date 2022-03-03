package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func main() {

	//nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	//singleNonDuplicate(nums)

	// 爬楼梯
	//n := 1
	//climbStairs(n)

	// 1791. 找出星型图的中心节点
	//edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	//findCenter(edges)

	// 917
	//s := "a-bC-dEf-ghIj"
	//reverseOnlyLetters(s)

	// 88
	//nums1 := []int {1,2,3,4}
	//nums2 := []int {1,1,1,1}
	//k := 3
	//merge(nums1,k,nums2,k)

	// 292
	//n := 1
	//canWinNim(n)

	// 258
	num := 258
	addDigits(num)
}

func singleNonDuplicate(nums []int) int {
	var temp = 0
	for i := 0; i <= len(nums)-1; i++ {
		if i != len(nums)-1 && nums[i] == nums[i+1] {
			continue
		}
		if nums[i] == nums[i-1] && i > 0 {
			continue
		}
		temp = nums[i]
	}

	return temp
}

func climbStairs(n int) int {
	// 这个递归超时了
	// 只有1、2两种的时候
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	//// 如3个阶梯 拆分为 1 + 2 ，以此类推..
	//return climbStairs(n - 1) + climbStairs(n - 2)

	a := 1
	b := 2
	temp := 0
	for i := 3; i <= n; i++ {
		temp = a
		a = b
		b = temp + b
	}

	return b
}

func findCenter(edges [][]int) int {

	a := edges[0][0]
	b := edges[0][1]
	if a == edges[1][0] || a == edges[1][1] {
		return a
	}

	return b
}

func reverseOnlyLetters(s string) string {

	ans := []byte(s)
	left, right := 0, len(s)-1
	for {
		for left < right && !unicode.IsLetter(rune(s[left])) { // 判断左边是否扫描到字母
			left++
		}
		for right > left && !unicode.IsLetter(rune(s[right])) { // 判断右边是否扫描到字母
			right--
		}
		if left >= right {
			break
		}
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	return string(ans)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	length1 := len(nums1)
	if m+n != length1 {

	}

	// 等于的时候直接排序num1
	if m < length1 {
		nums1 = append(nums1[0:m], nums2[0:n]...)
	}

	sort.Ints(nums1)
	fmt.Println(nums1)
}

func canWinNim(n int) bool {
	if n <= 3 {
		return true
	}

	if n%4 != 0 {
		return true
	}

	return false
}

func addDigits(num int) int {

	for {
		if num <= 9 {
			break
		}
		num = sumNum(num)
	}
	return num
}

func sumNum(num int) int {

	numVar := strconv.Itoa(num)
	total := 0
	for i := 0; i < len(numVar); i++ {
		to, _ := strconv.Atoi(string(numVar[i]))
		total += to
	}

	return total
}
