package leetcode

import (
	"sort"
)

func checkScramble(s1 string, s2 string, mp map[string]bool) bool {
	n := len(s1)
	s2_len := len(s2)
	if n != s2_len {
		return false
	}
	if s1 == s2 {
		return true
	}
	if n == 1 {
		return false
	}

	sort_s1 := []rune(s1)
	sort_s2 := []rune(s2)
	sort.Slice(sort_s1, func(i int, j int) bool { return sort_s1[i] < sort_s1[j] })
	sort.Slice(sort_s2, func(i int, j int) bool { return sort_s2[i] < sort_s2[j] })

	for i := 0; i < n; i++ {
		if sort_s1[i] != sort_s2[i] {
			return false
		}
	}

	key := s1 + "" + s2
	flag := false

	if val, isExist := mp[key]; isExist {
		return val
	}

	for i := 1; i < n; i++ {
		if checkScramble(s1[0:i], s2[0:i], mp) && checkScramble(s1[i:n], s2[i:n], mp) {
			flag = true
			return true
		}
		if checkScramble(s1[0:i], s2[n-i:n], mp) && checkScramble(s1[i:n], s2[0:n-i], mp) {
			flag = true
			return true
		}
		mp[key] = flag
	}
	return false
}

func isScramble(s1 string, s2 string) bool {
	mp := map[string]bool{}
	return checkScramble(s1, s2, mp)
}
