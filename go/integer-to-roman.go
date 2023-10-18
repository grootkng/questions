package main

import "fmt"

func intToRoman(num int) string {
	v := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	s := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""

	for i := 0; num > 0; i++ {
		for num >= v[i] {
			num -= v[i]
			res += s[i]
		}
	}

	return res
}

func main() {
	fmt.Println(intToRoman(50))
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
