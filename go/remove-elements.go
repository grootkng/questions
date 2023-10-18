package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = 999
			c++
		}
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return len(nums) - c
}

func main() {
	fmt.Println(removeElement([]int{4, 3, 2, 2, 1}, 3))          // 4
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))             // 2
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)) // 5
}
