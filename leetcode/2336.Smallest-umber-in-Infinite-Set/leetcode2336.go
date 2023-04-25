package leetcode

type SmallestInfiniteSet struct {
	set []int
}

func Constructor() SmallestInfiniteSet {
	set := make([]int, 1001)
	for i := 1; i <= 1000; i++ {
		set[i] = 1
	}
	return SmallestInfiniteSet{set: set}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	for i := 1; i <= 1000; i++ {
		if this.set[i] == 1 {
			this.set[i] = 0
			return i
		}
	}
	return 0
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num <= 1000 {
		this.set[num] = 1
	}
}
