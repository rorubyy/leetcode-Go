package leetcode

func validateStackSequences(pushed []int, popped []int) bool {
	sta := make([]int, 0)
	i := 0
	for _, val := range pushed {
		sta = append(sta, val)
		for len(sta) > 0 && popped[i] == sta[len(sta)-1] {
			sta = sta[:len(sta)-1]
			i++
		}
	}
	return len(sta)==0
}
