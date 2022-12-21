package main

import (
	"fmt"
	"github.com/wxnacy/wgo/arrays"
)

func main() {
	//strings.Replace()
	nums := []int{1,2,3,5,51,9,10,11,15,17,18,19,20,20,22,23,23,24,26,26,27,27,29,29,30,30,30,33,33,33,33,34,38,40,40,42,48,50,51,51,51,52,53,54,56,56,60,61,68,69,69,75,31,80,81,81,84,84,88,89,90,91,93,94,94,96,98,98,99,99}
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
