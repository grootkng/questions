package main

import "log"

func fib(n int) int {
	if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	}

	arr := []int{0, 1}
	for i := len(arr); i <= n; i++ {
		arr = append(arr, (arr[i-1] + arr[i-2]))
	}

	return arr[len(arr)-1]
}

func main() {
	log.Fatal(fib(4))
}
