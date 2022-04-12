package i32

/*
	最长有效括号，优先考虑dp
	f[i]: 以下标i结束的最长有效括号的长度
	没法进行空间优化
*/

func LongestValidParentheses(s string) int {
	f := make([]int, len(s))
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			continue
		}
		if i-1 >= 0 && s[i-1] == '(' {
			if i-2 >= 0 {
				f[i] = f[i-2] + 2
			} else {
				f[i] = 2
			}
		} else if i-1 >= 0 && i-f[i-1]-1 >= 0 && s[i-f[i-1]-1] == '(' {
			if i-f[i-1]-2 < 0 {
				f[i] = f[i-1] + 2
			} else {
				f[i] = f[i-1] + 2 + f[i-f[i-1]-2]
			}
		}
		cnt = max(cnt, f[i])
	}
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
