package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := len(nums) - 1; j > i; j-- {
// 			sum := nums[i] + nums[j]
// 			if sum == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		// 2
		temp := target - nums[i]
		for j := len(nums) - 1; j > i; j-- {
			if temp == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoNum := twoSum(nums, target)
	fmt.Println("the two num is ", twoNum)
	http.ListenAndServe("127.0.0.1:8088", nil)
}
