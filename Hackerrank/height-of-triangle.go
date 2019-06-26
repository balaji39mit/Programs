package hackerrank

import (
	"fmt"
	"math"
)

func height() {
	var base, area int64
	_, _ = fmt.Scanf("%d %d", &base, &area)
	height := int64(math.Ceil(float64(area) / float64(base) * 2))
	fmt.Println(height)
}
