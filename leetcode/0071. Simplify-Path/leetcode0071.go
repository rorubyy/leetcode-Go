package leetcode

import "strings"

func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	dir := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		if cur == ".." {
			if len(dir) != 0 {
				dir = dir[:len(dir)-1]
			}
		} else {
			if cur != "." && len(cur) > 0 {
				dir = append(dir, cur)
			}
		}
	}
	if len(dir) == 0 {
		return "/"
	}
	ans := strings.Join(dir, "/")
	ans = "/" + ans
	return ans

}
