package main

import (
	"fmt"
	"time"
)

func palindromic(key int) bool {
	fmt.Println("key", key)
	if key < 0 {
		return false
	}
	revertNumber := 0
	for key > revertNumber {
		revertNumber = revertNumber*10 + key%10
		key = key / 10
	}
	return revertNumber == key || key == revertNumber/10
}

func main() {
	key := 121
	start := time.Now()
	fmt.Println("start time:", start)
	isBool := palindromic(key)
	fmt.Println("end time:", time.Now(), time.Since(start))
	fmt.Println("debug:", isBool)
}
