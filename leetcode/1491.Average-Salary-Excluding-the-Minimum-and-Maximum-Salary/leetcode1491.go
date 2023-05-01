package leetcode

import (
	"math"
)

func average(salary []int) float64 {
	minSal, maxSal := math.MaxInt32, 0
	var ans float64 = 0
	for _, val := range salary {
		if val < int(minSal) {
			minSal = val
		}
		if val > maxSal {
			maxSal = val
		}
		ans += float64(val)
	}
	return (ans - float64(maxSal) - float64(minSal)) / float64(len(salary)-2)
}
