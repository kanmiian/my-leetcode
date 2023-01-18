package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
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
	//n := 4
	//fmt.Println(reinitializePermutation(n))
	//s := "ilovecodingonleetcode"
	//target := "code"
	//fmt.Println(rearrangeCharacters(s,target))

	sentence1 := ""
	sentence2 := ""
	fmt.Println(areSentencesSimilar(sentence1,sentence2))
}

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	strArr1 := strings.Fields(sentence1)
	strArr2 := strings.Fields(sentence2)
	if len(strArr2) > len(strArr1) {
		return  false
	}

	if len(strArr2) == len(strArr1) {
		return sentence1 == sentence2
	}

	ans := []int {}
	for i,v := range strArr1 {
		if strArr2[i] == v {
			ans = append(ans, i)
		}
	}
	if len(ans) < 1 {
		return false
	}
	if len(ans) != len(strArr2) {
		return  false
	}

	judgeNumsContinue(ans, len(strArr1) - 1)

	return true
}

func judgeNumsContinue(nums []int, max int) bool{
	if len(nums) == 1 {
		return nums[0] == 0 || nums[0] == max
	}

	if nums[0] != max - 1 {
		for i := 1; i < max; i++ {
			if nums[i] - 1 != nums[i - 1] {
				return false
			}
		}
	} else {
		for i := max; i > 0; i-- {
			if nums[i] - 1 != nums[i - 1] {
				return false
			}
		}
	}

	return true
}

func rearrangeCharacters(s string, target string) int {
	var targetMap, cntT [26]int
	for _, ch := range s { targetMap[ch-'a']++ }
	for _, ch := range target { cntT[ch-'a']++ }

	ans := len(s)
	for i, c := range targetMap {
		if cntT[i] > 0 {
			ans = min(ans, c/cntT[i])
		}
	}
	//fmt.Println(targetMap)
	return ans
}

func min(a, b int) int { if a > b { return b }; return a }

func reinitializePermutation(n int) int {
	target := make([]int, n)
	for i := range target {
		target[i] = i
	}
	perm := append([]int(nil), target...)
	fmt.Println(perm,"perm")
	step := 0
	for {
		step ++
		arr := make([]int,n)
		for i := range arr {
			if i%2 == 0 {
				arr[i] = perm[i / 2]
			} else {
				arr[i] = perm[n/2+i/2]
			}
		}
		perm = arr
		if equalArr(target,perm) {
			return step
		}
	}
}

func equalArr(a, b []int) bool {
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}

/**
	180. 统计各位数字之和为偶数的整数个数
 */
func countEven(num int) int {
	if num < 2 {
		return 0
	}
	ans := 0
	for i := 1; i <= num; i++ {
		sum := 0
		for x := i; x > 0; x /= 10 {
			sum += x % 10
		}
		if sum%2 == 0 {
			ans++
		}
	}

	return  ans
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
