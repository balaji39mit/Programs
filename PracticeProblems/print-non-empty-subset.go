package practice_problems

import (
	"fmt"
	"math"
	"sort"
)

func permute(str []rune) {
	output := make([]string, 0)
	nos := int(math.Pow(float64(2), float64(len(str))) - 1)
	for count := 0; count <= nos; count++ {
		s := make([]rune, 0)
		for i := 0; i < len(str); i++ {
			if count&(1<<uint32(i)) != 0 {
				s = append(s, str[i])
			}
		}
		output = append(output, string(s))
	}

	sort.Slice(output, func(i, j int) bool {
		if len(output[i]) == len(output[j]) {
			return output[i] < output[j]
		}
		return len(output[i]) < len(output[j])
	})
	for _, val := range output {
		fmt.Println(val)
	}

}
