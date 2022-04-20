package main

import "fmt"

func main() {
	n := 5
	a := []int{4, 2, 1, 3, 1}
	memo := map[int]int{}
	st := map[int]bool{}
	var f func(i int) int
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
		st[i] = true
		res := 0x3f3f3f3f
		for j := 0; j < i; j++ {
			if !st[j] {
				res = min(res, f(j)+1)
			}
		}
		if i+a[i] < n {
			if !st[i+a[i]] {
				res = min(res, f(i+a[i])+1)
			}
		}
		st[i] = false
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
