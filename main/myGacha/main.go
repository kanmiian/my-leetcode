package main

import (
	"fmt"
	"github.com/wxnacy/wgo/arrays"
)

func main() {
	//strings.Replace()
	nums := []int{23, 93, 88, 58, 46, 93, 87, 31, 82, 47}
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
