package main

import (
	"fmt"
)

func main() {
	str := "hello\000balaji"
	//	fmt.Println([]byte(str))
	for index, _ := range str {
		fmt.Println(str[index])

	}
	fmt.Println(len(str))
	/*subString := "\000"
	fmt.Println([]byte(subString))
	fmt.Println(len(subString))
	index := strings.Index(str, subString)
	fmt.Println(index)*/
}
