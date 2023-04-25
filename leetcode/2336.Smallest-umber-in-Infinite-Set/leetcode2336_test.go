package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem2336(t *testing.T) {
	fmt.Printf("-----Leetcode Problem 2336-----\n")
	obj := Constructor()
	param_1 := obj.PopSmallest()
	obj.AddBack(1)
	param_6 := obj.PopSmallest()
	param_2 := obj.PopSmallest()
	param_3 := obj.PopSmallest()
	obj.AddBack(2)
	obj.AddBack(3)
	param_4 := obj.PopSmallest()
	param_5 := obj.PopSmallest()
	fmt.Printf("%d,%d,%d,%d,%d,%d", param_1, param_6, param_2, param_3, param_4, param_5)

}
