package main

import (
	"fmt"
	"sort"
)

/**
贪心算法
*/
func main() {

	a := 1
	b := 8
	c := 8
	fmt.Println(maximumScore(a,b,c))
}

/**
  1753. 移除石子的最大得分
*/
func maximumScore(a int, b int, c int) int {
	total := a + b + c
	count := 0

	nums := []int{a, b, c}
	for i := 0; i < total; i++ {
		sort.Ints(nums)
		nums[1] --
		nums[2] --
		count ++
		if countZero(nums) {
			return count
		}
	}

	return count
}

func countZero(nums [] int)  bool{
	count := 0
	for _, v := range nums {
		if v ==0 {
			count++
		}
	}

	return count >= 2
}



















