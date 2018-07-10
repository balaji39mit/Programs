package main

import "fmt"

func romanToInt(s string) int {
	romanLiterals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	number := 0
	previousValue := 0
	for i := len(s) - 1; i >= 0; i-- {
		value := byte(s[i])
		currentValue := romanLiterals[value]
		if previousValue > currentValue {
			number = number + (currentValue * -1)
		} else {
			number = number + currentValue
		}
		previousValue = currentValue
	}
	return number
}

func main() {
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("XX"))
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
