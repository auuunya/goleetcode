package main

import (
	"fmt"
	"testing"
)

var nums = []int{2, 7, 11, 15, 12, 0, 8, 9, -12, 50, -48, -41, 3, 6, 1}
var target = 9

func BenchmarkTwoSum(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoNum := twoSum(nums, target)
		fmt.Println("two num is: ", twoNum)
	}
}
