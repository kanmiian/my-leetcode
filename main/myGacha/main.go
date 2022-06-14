package main

import (
	"fmt"
	"github.com/wxnacy/wgo/arrays"
)

func main() {
	//strings.Replace()
	nums := []int{69,9,9,12,13,14,14,15,16,21,23,24,27,30,31,32,34,35,35,39,47,48,49,51,51,55,55,58,60,66,68,68,69,71,72,72,73,75,81,82,83,87,88,88,90,93,94,95,95,99}
	countMoney(nums)
}

func countMoney(nums []int) int {

	total := 0
	arr := []int64{1, 7, 77, 100}
	in := 0
	for i := 0; i < len(nums)-1; i++ {
		inArr := arrays.ContainsInt(arr, int64(nums[i]))
		if inArr > 0 {
			in++
			continue
		}
		if nums[i] <= 10 {
			total += 15
		} else if nums[i] <= 20 {
			total += 5
		} else if nums[i] <= 30 {
			total += 1
		} else {
			total += 10
		}
	}

	fmt.Println(in)
	fmt.Println(total)
	return total
}
