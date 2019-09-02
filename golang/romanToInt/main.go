package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	//romanMap := make(map[string]int)
	romanMap := map[string]int{"I": 1, "IV": 3, "V": 5, "IX": 8, "X": 10,
		"XL": 30, "L": 50, "XC": 80, "C": 100, "CD": 300, "D": 500, "CM": 800, "M": 1000}
	//fmt.Println(romanMap)
	var num int
	x := 0
	for {

		if a, ok := romanMap[string(s[x])]; ok {
			num = num + a
		}
		if x == len(s)-1 {
			break
		}
		//fmt.Println(string(s[x : x+2]))
		if a, ok := romanMap[string(s[x:x+2])]; ok {
			x++
			num = num + a
			//fmt.Println(a)
		}
		x++
		if x == len(s) {
			break
		}
	}
	return num
}
