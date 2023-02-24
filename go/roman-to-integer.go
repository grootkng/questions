package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	res := 0
	cs := s
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	if strings.Contains(cs, "IV") {
		res = m["V"] - 1
		cs = strings.Replace(cs, "IV", "", -1)
	}

	if strings.Contains(cs, "IX") {
		res = res + (m["X"] - 1)
		cs = strings.Replace(cs, "IX", "", -1)
	}

	if strings.Contains(cs, "XL") {
		res = res + (m["L"] - 10)
		cs = strings.Replace(cs, "XL", "", -1)
	}

	if strings.Contains(cs, "XC") {
		res = res + (m["C"] - 10)
		cs = strings.Replace(cs, "XC", "", -1)
	}

	if strings.Contains(cs, "CD") {
		res = res + (m["D"] - 100)
		cs = strings.Replace(cs, "CD", "", -1)
	}

	if strings.Contains(cs, "CM") {
		res = res + (m["M"] - 100)
		cs = strings.Replace(cs, "CM", "", -1)
	}

	for i := 0; i < len(cs); i++ {
		res = res + m[string(cs[i])]
	}

	return res
}

func main() {
	fmt.Println(romanToInt("III"))     // 3 OK
	fmt.Println(romanToInt("IV"))      // 4 OK
	fmt.Println(romanToInt("LVIII"))   // 58 OK
	fmt.Println(romanToInt("MCMXCIV")) // 1994 OK
}
