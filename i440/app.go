package i440

func findKthNumber(n int, k int) int {
	res := 0
	i := 0    // 当前位置
	lcnt := 0 // 左边的元素个数
	for lcnt < k {
		curCnt := getCount(n, i)
		if curCnt+lcnt <= k {
			lcnt += curCnt
			i++
			res++
		} else {
			lcnt++
			i = i * 10
			res *= 10
		}
	}
	return res
}

/*
   获取[0, n]中前缀为k的数字个数
   通过这个函数可以求出每个节点的元素个数
*/
func getCount(n int, k int) int {
	if k == 0 {
		return 1
	}
	cnt := 0

	x := k
	base := 1
	for x <= n {
		if x+base <= n {
			cnt += base
		} else {
			cnt += n - x + 1
		}
		x *= 10
		base *= 10
	}

	return cnt
}
