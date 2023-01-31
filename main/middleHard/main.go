package main

import (
	"fmt"
	"my-leetcode/main/common"
	"sort"
	"strconv"
	"strings"
)
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := common.NewLinkedList([]int {1,2,3,4,5,6,7})
	a := 3
	b := 4
	list2 := common.NewLinkedList([]int {1000000,1000001,100000})
	fmt.Println(mergeInBetween(list1,a,b,list2))
}

func mergeInBetween(list1 *common.ListNode, a int, b int, list2 *common.ListNode) *common.ListNode {
	tmpA := list1
	for i := 0;i < a - 1; i ++ {
		tmpA = tmpA.Next
	}
	tmpB := tmpA
	for i := 0; i < b - a + 2; i ++ {
		tmpB = tmpB.Next
	}
	tmpA.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = tmpB
	return  list1
}


func countNicePairs(nums []int) int {
	ans := 0
	//nums[i]-rev(nums[i])=nums[j]-rev(nums[j])
	//numMaps := map[int]int{}
	//for i := 0; i <= len(nums) - 1; i ++ {
	//	for j := len(nums) - 1; j > 0; j -- {
	//		if j == i {
	//			continue
	//		}
	//		if nums[i] - reverse(nums[i]) == nums[j] - reverse(nums[j]) {
	//			fmt.Println(i,"iiii",j,"jjjjjj")
	//			if v, state := numMaps[min(i,j)]; state {
	//				if v == max(i,j) {
	//					continue
	//				}
	//			}
	//			ans ++
	//			numMaps[min(i,j)] = max(i,j)
	//			fmt.Println(numMaps,"numsMapppp")
	//		}
	//	}
	//}
	rev := func(x int) (y int) {
		for ; x > 0; x /= 10 {
			y = y*10 + x%10
		}
		return
	}
	cnt := map[int]int{}
	const mod int = 1e9 + 7
	for _, x := range nums {
		y := x - rev(x)
		ans = (ans + cnt[y]) % mod
		cnt[y]++
	}

	return ans
}

func reverse(x int, y int) int {
	for ; x > 0; x /= 10 {
		y = y*10 + x%10
	}
	return y
}

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	strArr1 := strings.Fields(sentence1)
	strArr2 := strings.Fields(sentence2)

	i, n := 0, len(strArr1)
	j, m := 0, len(strArr2)
	for i < n && i < m && strArr1[i] == strArr2[i] {
		i++
	}
	for j < n-i && j < m-i && strArr1[n-j-1] == strArr2[m-j-1] {
		j++
	}
	//return i+j == getMid(n, m)
	return false
}

func areSentencesSimilar2(sentence1 string, sentence2 string) bool {
	strArr1 := strings.Fields(sentence1)
	strArr2 := strings.Fields(sentence2)

	if len(strArr2) == len(strArr1) {
		return sentence1 == sentence2
	}

	var ans []int
	longArr := strArr1
	shortArr := strArr2
	if len(strArr2) > len(strArr1) {
		longArr = strArr2
		shortArr = strArr1
	}
	// todo
	for i, v := range longArr {
		for i2, _ := range shortArr {
			if shortArr[i2] == v {
				ans = append(ans, i)
			}
		}
	}
	if len(ans) < 1 {
		fmt.Println(ans, "ans")
		return false
	}
	if len(ans) != len(shortArr) {
		fmt.Println(ans, "长度不相等")
		return false
	}

	return judgeNumsContinue(ans, len(longArr)-1)
}

func judgeNumsContinue(nums []int, max int) bool {
	if len(nums) == 1 {
		fmt.Println(nums, max, "长度等于1", nums[0] == 0 || nums[0] == max)
		return nums[0] == 0 || nums[0] == max
	}

	if nums[0] == 0 {
		for i := 1; i < max; i++ {
			if nums[i]-1 != nums[i-1] {
				fmt.Println(nums, "包含了首部", max)
				return false
			}
		}
	}

	if nums[len(nums)-1] == max-1 {
		for i := max; i > 0; i-- {
			if nums[i]-1 != nums[i-1] {
				fmt.Println(nums, "包含了尾部", max)
				return false
			}
		}
	}

	// 只有中间 肯定是错的  刚好是收尾两个直接返回true

	return true
}

/**

 */
func evaluate(s string, knowledge [][]string) string {
	undefineVal := "?"
	stringMap := map[string]string{}
	ans := strings.Builder{}
	start := -1
	for _, arr := range knowledge {
		stringMap[arr[0]] = arr[1]
	}
	for i, c := range s {
		if c == '(' {
			start = i
		} else if c == ')' {
			if t, state := stringMap[s[start+1:i]]; state {
				ans.WriteString(t)
			} else {
				ans.WriteString(undefineVal)
			}
			start = -1
		} else if start < 0 {
			ans.WriteString(string(c))
		}
	}

	return ans.String()
}

/**
1803. 统计异或值在范围内的数对有多少
*/
func countPairs(nums []int, low int, high int) int {
	ans := 0
	sort.Ints(nums)
	i, j := 0, 0
	for i = 0; i < len(nums); i++ {
		for j = i + 1; j < len(nums); j++ {
			xor := nums[i] ^ nums[j]
			if xor >= low && xor <= high {
				ans++
			}
		}
	}
	return ans

	//n := len(nums)
	//res := 0
	//sort.Ints(nums)
	//for i := 0; i < n-1; i++ {
	//	a := nums[i]
	//	for j := i + 1; j < n; j++ {
	//		xor := a ^ nums[j]
	//		if xor >= low && xor <= high {
	//			res++
	//		}
	//	}
	//}
	//return res
}

/**
2042. 检查句子中的数字是否递增
*/
func areNumbersAscending(s string) bool {
	strArray := strings.Fields(s)
	pre := 0
	for _, v := range strArray {
		if common.IsNum(v) {
			fmt.Println(v)
			nextNum, _ := strconv.Atoi(v)
			if nextNum <= pre {
				return false
			}
			pre, _ = strconv.Atoi(v)
		}
	}
	return true
}

func maxArea(height []int) int {

	length := len(height)
	start := 0
	end := length - 1
	total := 0
	min := common.GetMin(height[start], height[end])
	total = min * (end - start)
	for {
		if start == end {
			break
		}

		if height[start] < height[end] {
			start++
		} else {
			end--
		}
		temp := common.GetMin(height[start], height[end])
		temoTotal := temp * (end - start)
		if temoTotal > total {
			total = temoTotal
		}
	}
	fmt.Println(total, start, end)
	return total
}

func isValidSudoku(board [][]byte) bool {

	for i, boa := range board {
		fmt.Println(i)
		fmt.Println(boa)
	}

	return true
}

func minEatingSpeed(piles []int, h int) int {
	arrLen := len(piles)
	sort.Ints(piles)
	left := 0
	right := piles[arrLen-1]
	// 个数相同，以最大值为结果
	if arrLen == h {
		return piles[arrLen-1]
	}

	for left < right {
		mid := (right-left)/2 + left

		if getMid(piles, mid) <= h {
			right = mid
		} else if getMid(piles, mid) > h {
			left = mid + 1
			fmt.Println(left, mid+1)
		}

	}

	fmt.Println(left)
	return left
}

func getMid(piles []int, x int) int {
	hours := 0
	for _, pile := range piles {
		hours += pile / x
		if pile%x > 0 {
			hours++
		}
	}
	return hours
}



/**
插入元素到数组内
*/
func insertArr(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
