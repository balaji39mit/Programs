package main

func reverse(x int) int {
	var output int
	for x != 0 {
		reminder := x % 10
		output = output*10 + reminder
		x = x / 10
	}
	return output
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := reverse(x)
	return x == y
}

func main() {
	print(isPalindrome(121))
}
