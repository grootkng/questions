package main

import "log"

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}

	arr := []int{1, 1}
	for i := 2; i <= n; i++ {
		arr = append(arr, (arr[i-1] + arr[i-2]))
	}

	return arr[n]
}

func main() {
	log.Fatal(climbStairs(4))
}
