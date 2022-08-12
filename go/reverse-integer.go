package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func solution(x int) int {
	var sg int = 1
	if x < 0 {
		sg = -1
	}

	n := x * sg
	str := strconv.Itoa(n)
	spStr := strings.SplitAfter(str, "")

	tSw := len(spStr) - 1
	for i := 0; i < (len(str) / 2); i++ {
		spStr[i], spStr[tSw] = spStr[tSw], spStr[i]
		tSw--
	}

	sj := strings.Join(spStr, "")
	cvN, _ := strconv.Atoi(sj)

	if x < 0 {
		cvN = cvN * sg
	}

	if cvN >= math.MaxInt32 || cvN <= math.MinInt32 {
		return 0
	}

	return cvN
}

func main() {
	fmt.Printf("Solution: %v", solution(901000))
	fmt.Println()

	fmt.Printf("Solution: %v", solution(-901000))
	fmt.Println()

	fmt.Printf("Solution: %v", solution(123))
	fmt.Println()

	fmt.Printf("Solution: %v", solution(-120))
	fmt.Println()

	return
}
