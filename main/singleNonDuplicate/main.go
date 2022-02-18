package main

func main() {

	//nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	//singleNonDuplicate(nums)

	// 爬楼梯
	//n := 1
	//climbStairs(n)

	// 1791. 找出星型图的中心节点
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	findCenter(edges)
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
