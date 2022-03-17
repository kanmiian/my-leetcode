package main

import (
	"fmt"
)

func main() {

	// 11. 盛最多水的容器
	//height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//maxArea(height)

	// 有效的数独
	//board := [][]btye
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	isValidSudoku(board)
}

func maxArea(height []int) int {

	length := len(height)
	start := 0
	end := length - 1
	total := 0
	min := getMin(height[start], height[end])
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
		temp := getMin(height[start], height[end])
		temoTotal := temp * (end - start)
		if temoTotal > total {
			total = temoTotal
		}
	}
	fmt.Println(total, start, end)
	return total
}

func isValidSudoku(board [][]byte) bool {

	for i,boa := range board {
		fmt.Println(i)
		fmt.Println(boa)
	}

	return true
}













/**
获取较小值
*/
func getMin(start int, end int) int {
	if start > end {
		return end
	} else {
		return start
	}
}

/**
获取较大值
*/
func getMax(start int, end int) int {
	if start > end {
		return start
	} else {
		return end
	}
}
