package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	for i := 0; ; i++ {
		for _, str := range strs {
			if i == len(str) || str[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
}

func main() {
	s := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(s))

	s = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(s))
}
