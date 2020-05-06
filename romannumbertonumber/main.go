package main

import "fmt"

func romannumbertonumber(s string) int {
	hashMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	number := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && hashMap[string(s[i])] < hashMap[string(s[i+1])] {
			number -= hashMap[string(s[i])]
		} else {
			number += hashMap[string(s[i])]
		}
	}
	return number
}
func main() {
	sum := romannumbertonumber("III")
	fmt.Println("sum is: ", sum)
}
