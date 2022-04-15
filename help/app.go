package main

import "fmt"

func main() {
	n := 5
	a := []int{4, 2, 1, 3, 1}
	memo := map[int]int{}
	var f func(i int) int
	st := map[int]bool{}
	f = func(i int) int {
		if v, ok := memo[i]; ok {
			return v
		}
		if i >= n {
			memo[i] = 0
			return 0
		}
		if i+a[i] >= n {
			memo[i] = 1
			return 1
		}
		res := 0x3f3f3f3f
		for j := 0; j < i; j++ {
			if !st[j] {
				st[j] = true
				res = min(res, f(j)+1)
				st[j] = false
			}
		}
		if i+a[i] < n {
			st[i+a[i]] = true
			res = min(res, f(i+a[i])+1)
			st[i+a[i]] = false
		}
		memo[i] = res
		return res
	}
	fmt.Println(f(0))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
