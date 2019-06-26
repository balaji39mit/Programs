package hackerrank

import (
	"bufio"
	"io"
	"math"
	"strings"
)

/*
 * Complete the primeCount function below.
 */
func primeCount(n int64) int32 {
	count := int32(0)
	if n <= 1 {
		return count
	}

	if n%2 == 0 {
		count++
	}
	for (n % 2) == 0 {
		n = n / 2
	}

	for i := int64(3); i <= int64(math.Sqrt(float64(n))); i = i + 2 {
		if n%i == 0 {
			count++
			for n%i == 0 {
				n = n / i
			}
		}
	}

	if n > 1 {
		count = 1
	}

	return count
}

func checkError1(err error) {
	if err != nil {
		panic(err)
	}
}

func readLine1(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
