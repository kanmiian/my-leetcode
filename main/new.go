package main

import (
	"fmt"
	"math"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 重复元素②
	//nums := []int{1,2,3,3,1,2}
	//containsNearbyDuplicate(nums, 3)

	// 53. 最大子数组和
	//nums := []int{1}
	//fmt.Println(maxSubArray(nums))

	// 1984. 学生分数的最小差值
	//nums := []int{90}
	//k := 1
	//minimumDifference(nums, k)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	newArr := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		newArr[nums[i]] = append(newArr[nums[i]], i)
	}

	for _, value := range newArr {
		if len(value) > 1 {
			for i := 0; i < len(value)-1; i++ {
				if value[i+1]-value[i] <= k {
					return true
				}
			}
		}
	}
	return false
}

func maxSubArray(nums []int) int {
	var sum int
	sum = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > sum {
			sum = nums[i]
		}
	}
	return sum
}

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)

	diff := math.MaxInt32
	fmt.Println(len(nums) - k + 1)
	for i := range nums[:len(nums)-k+1] {
		if (nums[i+k-1] - nums[i]) < diff {
			diff = nums[i+k-1] - nums[i]
		}
	}

	return diff
}
