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
	//words := []string{"a","banana","app","appl","ap","apply","apple"}
	//longestWord(words)

	// 交替位二进制
	//n := 5
	//hasAlternatingBits(n)

	// 多少个1
	//var n uint32
	//n = 00000000000000000000000000001011
	//hammingWeight(n)

	// 提莫攻击
	//timeSeries := []int {1,2}
	//duration := 2
	//findPoisonedDuration(timeSeries, duration)

	// 校验ip地址
	//queryIP := "172.16.254.1"
	//validIPAddress(queryIP)

	// 小镇法官
	//trust := [][]int{{1, 2}}
	//n := 2
	//findJudge(n, trust)

	// 硬币排列
	//n := 15
	//arrangeCoins(n)

	// 有效的回旋镖
	//points := [][]int{{1, 1}, {2, 2}, {3, 3}}
	//isBoomerang(points)

	// 1051. 高度检查器
	//heights := []int{1, 1, 4, 2, 1, 3}
	//heightChecker(heights)

}




func duplicateZeros(arr []int) {
	//length := len(arr)
	// 得到的res是最后的长度
	//res := 0
	//for i := 0; i < length; i++ {
	//	if arr[i] != 2 {
	//		arr[res] = arr[i]
	//		res++
	//	}
	//}
	//top, i := 0, -1
	//for top < len {
	//	i ++
	//	if arr[i] != 0 {
	//		top++
	//	} else {
	//		top += 2
	//	}
	//}
	//j := len -1
	//if top == len +1 {
	//	arr[j] = 0
	//	j --
	//	i --
	//}
	//for j >= 0 {
	//	arr[j] = arr[i]
	//	j--
	//	if arr[i] == 0 {
	//		arr[j] = arr[i]
	//		j--
	//	}
	//	i--
	//}
	fmt.Println(arr)
}

func heightChecker(heights []int) int {

	sortHeight := make([]int, len(heights))
	copy(sortHeight, heights)
	sort.Ints(heights)

	fmt.Println(sortHeight)
	fmt.Println(heights)
	count := 0
	for i, _ := range heights {
		if heights[i] != sortHeight[i] {
			count++
		}

	}

	return count
}

func isBoomerang(points [][]int) bool {

	// y = kx+b
	point1 := points[0]
	point2 := points[1]
	point3 := points[2]
	r1 := point1[0] - point2[0]
	c1 := point1[1] - point2[1]
	r2 := point2[0] - point3[0]
	c2 := point2[1] - point3[1]

	return r1*c2 != r2*c1
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
	for i := range str {
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

// todo
func hasAlternatingBits(n int) bool {

	//now := 0
	//pre := 2
	for pre := 2; n != 0; n /= 2 {
		now := n % 2
		if now == pre {
			return false
		}
		pre = now
	}

	return true
}

// todo
func hammingWeight(num uint32) int {

	result := 0
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			result++
		}
	}

	fmt.Println(result)
	return result
}

// 提莫攻击
func findPoisonedDuration(timeSeries []int, duration int) int {
	//我们只需要对数组进行一次扫描就可以计算出总的中毒持续时间。我们记录艾希恢复为未中毒的起始时间
	//expired，设艾希遭遇第 i 次的攻击的时间为 timeSeries[i]。当艾希遭遇第 i 攻击时：
	//
	//如果当前他正处于未中毒状态，则此时他的中毒持续时间应增加
	//duration，同时更新本次中毒结束时间expired
	//等于 {timeSeries}[i]+ \textit{duration}timeSeries[i]+duration​；
	//如果当前他正处于中毒状态，由于中毒状态不可叠加，
	//我们知道上次中毒后结束时间为 \textit{expired}expired​​，
	//本次中毒后结束时间为 \textit{timeSeries}[i] + \textit{duration}timeSeries[i]+duration​​，
	//因此本次中毒增加的持续中毒时间为 \textit{timeSeries}[i] + \textit{duration} -\textit{expired}timeSeries[i]+duration−expired​​；
	//我们将每次中毒后增加的持续中毒时间相加即为总的持续中毒时间。

	if duration == 0 {
		return 0
	}

	time := 0
	ans := 0
	for _, t := range timeSeries {
		if t >= time {
			ans += duration
		} else {
			ans += t + duration - time
		}
		time = t + duration
	}

	fmt.Println(time)
	return time
}

// 排列硬币
func arrangeCoins(n int) int {

	no := 0
	if n == 1 {
		return 1
	}
	for i := 1; i <= n; i++ {
		if n-i > 0 {
			no++
		}
		if n-i == i+1 {
			no++
			break
		}
		// 最后一条不符合长度 结束
		if n-i < i+1 {
			break
		}
		n -= i
	}

	fmt.Println(no)
	return no
}

// RemoveRepeatedElementAndEmpty 数组去重string
func removeRepeatElement(list []string) []string {
	// 创建一个临时map用来存储数组元素
	temp := make(map[string]bool)
	index := 0
	for _, v := range list {
		// 遍历数组元素，判断此元素是否已经存在map中
		_, ok := temp[v]
		if ok {
			list = append(list[:index], list[index+1:]...)
		} else {
			temp[v] = true
		}
		index++
	}
	return list
}

// RemoveRepeatedElementAndEmpty 数组去重int
func uniqueArr(arr []int) []int {
	newArr := make([]int, 0)
	tempArr := make(map[int]bool, len(newArr))
	for _, v := range arr {
		if tempArr[v] == false {
			tempArr[v] = true
			newArr = append(newArr, v)
		}
	}
	return newArr
}

// todo 验证IP地址  不知道怎么识别ipv6
func validIPAddress(queryIP string) string {
	//strArr := strings.Split(queryIP, ".")
	//IPv4 := "IPv4"
	//IPv4 := "IPv6"
	//Neither := "Neither"
	//if len(strArr) != 4 {
	//	return Neither
	//}
	//
	//

	return ""
}

// 小镇法官
func findJudge(n int, trust [][]int) int {
	judgeMap := map[int]int{}
	result := -1
	if n == 1 {
		return n
	}
	if n == 0 {
		return result
	}
	if len(trust) < 1 {
		return result
	}
	for _, value := range trust {
		judgeMap[value[1]]++
	}

	for key, value := range judgeMap {
		if value == n-1 {
			result = key
		}
	}
	for _, v := range trust {
		if v[0] == result {
			return -1
		}
	}
	fmt.Println(judgeMap)
	fmt.Println(result)
	return result
}

func insertArr(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}


