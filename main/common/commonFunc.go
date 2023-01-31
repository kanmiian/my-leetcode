package common

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}

func NewLinkedList(val []int) *ListNode {
	nodes := make([]*ListNode, len(val)+1)
	for i := len(val) - 1; i >= 0; i-- {
		nodes[i] = NewListNode(val[i], nodes[i+1])
	}

	return nodes[0]
}



/**
	数字对比等常用方法
 */

func GetMid(piles []int, x int) int {
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
获取较小值
*/
func GetMin(start int, end int) int {
	if start > end {
		return end
	} else {
		return start
	}
}

/**
获取较大值
*/
func GetMax(start int, end int) int {
	if start > end {
		return start
	} else {
		return end
	}
}

/**
插入元素到数组内
*/
func InsertArr(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

/**
	判断是否为数字
 */
func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
