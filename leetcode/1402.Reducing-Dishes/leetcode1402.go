package leetcode

import (
	"math"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	dishLen, nowValue := len(satisfaction), 0
	dp := make([]int, dishLen)
	sort.Sort(sort.Reverse(sort.IntSlice(satisfaction)))
	for i := 0; i < dishLen; i++ {
		nowValue = 0
		for j := i; j >= 0; j-- {
			nowValue += satisfaction[j] * (i - j + 1)
		}
		dp[i] = int(math.Max(float64(nowValue), float64(dp[int(math.Max(float64(i-1),0))])))
	}
	return dp[dishLen-1]
}
