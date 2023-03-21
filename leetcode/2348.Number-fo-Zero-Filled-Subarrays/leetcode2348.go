package leetcode

func zeroFilledSubarray(nums []int) int64 {
	var output, contiguousZero int64
	output = 0
	contiguousZero = 0
	for i := 0; i <= len(nums); i++ {
		if i < len(nums) && nums[i] == 0 {
			contiguousZero++
			continue
		}
		output += (1 + contiguousZero) * contiguousZero / 2
		contiguousZero = 0
	}
	return output
}
