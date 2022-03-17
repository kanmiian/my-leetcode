package main

import (
	"fmt"
	"math"
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
	//num := 258
	//addDigits(num)

	// 7进制
	//n := -7
	//convertToBase7(n)

	// 比较大小
	//a := 2
	//b := 5
	//maximum(a, b)

	// 两个列表最小索引总和
	//list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	//list2 := []string{"KFC", "Burger King", "Tapioca Express"}
	//findRestaurant(list1, list2)

	// 最佳售出时间
	//prices := []int{3, 2, 6, 5, 0, 3}
	//maxProfit(prices)

	// 是否回文数
	//x := 123
	//isPalindrome(x)

	// 词典中最长的单次
	words := []string{"a","banana","app","appl","ap","apply","apple"}
	longestWord(words)
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

func convertToBase7(num int) string {

	if num == 0 {
		return "0"
	}
	num = int(math.Abs(float64(num)))
	str := ""
	for {
		if num == 0 {
			break
		}
		str = strconv.Itoa(num%7) + str
		num /= 7
	}

	if num < 0 {
		str = "-" + str
	}

	return str
}

func maximum(a int, b int) int {

	//k := 0
	//result1 := a * k - b * (k - 1)

	// todo 这样计算其实感觉没啥意义...只不过这样不需要加入＞＜ 但其实==也是运算符，不应该使用的
	result := float64(a-b) + math.Abs(float64(a-b))

	if int(result) == 0 {
		return b
	}
	return a
}

func findRestaurant(list1 []string, list2 []string) []string {

	ans := []string{}
	//ansNo := []int{}
	min := math.MaxInt16
	for i, str := range list1 {
		for i2, str2 := range list2 {
			if str == str2 {
				if i+i2 == min {
					ans = append(ans, str)
				}
				if i+i2 < min {
					min = i + i2
					ans = []string{str}
				}
			}
		}
	}

	fmt.Println(ans)
	//if len(ans) == 1 {
	//	return ans
	//}
	//
	//fin := ansNo[0]
	//finNo := 0
	//findNo := []int{}
	//for i, no := range ansNo {
	//	if no < fin {
	//		fin = no
	//		finNo = i
	//	}
	//	if no == fin && i != 0 {
	//		findNo = append(findNo, i)
	//	}
	//}
	//
	//if len(findNo) > 1 {
	//	fmt.Println(findNo)
	//}
	//result := []string{}
	//fmt.Println(append(result, ans[finNo]))
	return nil
}

func longestWord(words []string) string {

	//newOrder := []int{}
	//for _, word := range words {
	//	newOrder = append(newOrder,len(word))
	//}
	//sort.Ints(newOrder)
	//newWords := []string{}
	//for _,order := range newOrder{
	//	newWords = append(newWords,words[order-1])
	//}
	//fmt.Println(newWords)
	// 使用sort.SLice进行排序
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) != len(words[j]) {
			return len(words[i]) < len(words[j])
		}
		return words[i] > words[j]
	})
	//fmt.Println(words)
	mp := map[string]struct{}{}
	mp[""] = struct{}{}
	res := ""
	for i := range words {
		if _, has := mp[words[i][:len(words[i])-1]]; has {
			mp[words[i]] = struct{}{}
			res = words[i]
		}
	}
	fmt.Println(res)
	return res
}

func maxProfit(prices []int) int {

	min := prices[0]
	max := prices[0]
	diff := 0
	for _, price := range prices {
		if price > max {
			max = price
			if diff < max-min {
				diff = max - min
			}
		}
		if price < min {
			min = price
			max = price
		}
	}

	fmt.Println(diff)
	fmt.Println(min)
	fmt.Println(max)
	return diff
}

func isPalindrome(x int) bool {

	if x >= 0 && x < 10 {
		return true
	}

	if x < 0 || x%10 == 0 {
		return false
	}

	str := strconv.Itoa(x)
	len := len(str)
	//fmt.Println(len, str)
	for i, _ := range str {
		if len-i-1 < 0 {
			break
		}
		fmt.Println(str[i], str[len-i-1])
		if str[i] != str[len-i-1] {
			return false
		}

	}

	fmt.Println(55555)
	return true
}
