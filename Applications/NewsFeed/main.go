package main

import "fmt"

func main() {
	fmt.Println("Entering into application. Choose one of this.")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Sprintf("Panic occurred..! Reason : %5s", r))
		}
	}()
	for {
		fmt.Println("post/follow/comment/vote/feed")
		var input string
		_, _ = fmt.Scanln(&input)
		if val, ok := features[input]; ok {
			val()
		} else {
			fmt.Println("Please provide the supported features")
		}
	}
}
