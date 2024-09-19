package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	currentCount := 0
	for _, num := range nums {
		if num == 1 {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 0
		}
	}
	return maxCount
}

func main() {
	nums := []int{1, 1, 0, 0, 1, 1, 1}
	fmt.Println("Max consecutive ones:", findMaxConsecutiveOnes(nums))
}
