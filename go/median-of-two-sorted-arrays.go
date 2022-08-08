package main

import (
	"fmt"
	"sort"
)

func solution1(nums1 []int, nums2 []int) float64 {
	mergedArr := append(nums1, nums2...)

	for i := 1; i < (len(nums1) + len(nums2)); i++ {
		j := i
		for j > 0 {
			if mergedArr[j-1] > mergedArr[j] {
				mergedArr[j-1], mergedArr[j] = mergedArr[j], mergedArr[j-1]
			}
			j -= 1
		}
	}

	if len(mergedArr)%2 == 0 {
		i := (len(mergedArr) - 1) / 2
		return (float64(mergedArr[i]) + float64(mergedArr[i+1])) / 2.0
	} else {
		return float64(mergedArr[int((len(mergedArr)-1)/2)])
	}
}

func solution2(nums1 []int, nums2 []int) float64 {
	mergedArr := append(nums1, nums2...)
	sort.Ints(mergedArr)

	if len(mergedArr)%2 == 0 {
		i := (len(mergedArr) - 1) / 2
		return (float64(mergedArr[i]) + float64(mergedArr[i+1])) / 2.0
	} else {
		return float64(mergedArr[int((len(mergedArr)-1)/2)])
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}

	str1 := fmt.Sprint("Solution 1: ", solution1(nums1, nums2))
	str2 := fmt.Sprint("Solution 2: ", solution2(nums1, nums2))

	fmt.Println(str1)
	fmt.Println(str2)
}
