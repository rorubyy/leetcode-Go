package leetcode

func maxElement(arr []int) int {
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxKid := maxElement(candies)
	output := make([]bool,len(candies))
	for i,v := range candies{
		if v+extraCandies>=maxKid{
			output[i]=true
		}else{
			output[i]=false
		}
	}
	return output
}
