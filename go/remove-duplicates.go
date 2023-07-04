package main

import "fmt"

func removeDuplicates(nums []int) int {
	r := []int{}
	seen := make(map[int]bool)

	for _, n := range nums {
		if !seen[n] {
			seen[n] = true
			r = append(r, n)
		}
	}

	for i := 0; i < len(r); i++ {
		fmt.Print(r[i])
	}

	return len(r)
}

func main() {
	inp := []int{1, 1, 2}
	fmt.Println(removeDuplicates(inp))

	inp = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(inp))

	inp = []int{1, 2, 3}
	fmt.Println(removeDuplicates(inp))

	inp = []int{1, 2}
	fmt.Println(removeDuplicates(inp))

	inp = []int{1}
	fmt.Println(removeDuplicates(inp))

	inp = []int{1, 1, 2, 2, 3, 3}
	fmt.Println(removeDuplicates(inp))

	inp = []int{0, 0, 0, 0, 0}
	fmt.Println(removeDuplicates(inp))
}
