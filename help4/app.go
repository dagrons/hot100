package main

import (
	"fmt"
	"sort"
)

func main() {
	var m int
	fmt.Scan(&m)
	for k := 0; k < m; k++ {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		a2 := make([]int, n)
		color := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
			a2[i] = a[i]
		}
		for i := 0; i < n; i++ {
			fmt.Scan(&color[i])
		}

		sort.Slice(a2, func(i, j int) bool {
			return a2[i] < a2[j]
		})

		pos2 := map[int]int{} // a2[i]在a2[]中的位置
		for i := 0; i < n; i++ {
			pos2[a2[i]] = i
		}

		pos := map[int]int{}
		for i := 0; i < n; i++ {
			pos[a[i]] = i
		}

		for i := 0; i < n; i++ { // 对每一个a[i]
			for j := 0; j < pos2[a[i]]; j++ {
				if pos[a[j]] > i && color[i] == color[pos[a2[j]]] {
					fmt.Println("No")
					return
				}
			}
		}

		fmt.Println("Yes")
	}
}
