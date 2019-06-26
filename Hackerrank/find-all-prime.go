package hackerrank

import (
	"bufio"
	"io"
	"strings"
)

/*
 * Complete the primeCount function below.
 */

func isPrime(n int64) bool {
	// Corner cases
	// This is checked so that we can skip
	// middle five numbers in below loop

	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := int64(5); i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}

	}
	return true
}

func CountPrime(n int64) int32 {
	if n <= 1 {
		return 0
	}
	count := int32(0)
	if n/2 >= 1 {
		count++
		n = n / 2
	}
	for i := int64(3); n/i >= 1; i = i + 2 {
		if isPrime(i) {
			count++
			n = n / i
		}
	}
	return count
}

func readLine2(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError2(err error) {
	if err != nil {
		panic(err)
	}
}
