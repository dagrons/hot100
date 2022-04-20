package main

import "fmt"

var cnt int // 情况个数

func dfs(t int, ans int, w int, N int) { // ans当前分数, w当前答错的个数

	if w <= 3 && ans == N { // 后面的题连续全部答错
		cnt++
		return
	}
	if w >= 3 || ans > N || t >= 25 {
		return
	}
	if t < 10 {
		dfs(t+1, ans+2, w, N)
		dfs(t+1, ans, w+1, N)
	} else if t < 20 {
		dfs(t+1, ans+4, w, N)
		dfs(t+1, ans, w+1, N)
	} else {
		dfs(t+1, ans+8, w, N)
		dfs(t+1, ans, w+1, N)
	}
}

func main() {
	dfs(0, 0, 0, 2)
	fmt.Println(cnt)
}
