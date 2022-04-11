package i5

func LongestPalindrome(s string) string {
	var res string
	n := len(s)
	cnt := 0

	for i := range s { // 偶数
		j := 0
		for ; i-j >= 0 && i+j < n && s[i-j] == s[i+j]; j++ {
		}
		if 2*j+1 > cnt {
			cnt = 2*j + 1
			res = s[i-j+1 : i+j]
		}
	}

	for i := 0; i < n-1; i++ { // 奇数
		j := 0
		for ; i-j >= 0 && i+1+j < n && s[i-j] == s[i+1+j]; j++ {
		}
		if 2*(j+1) > cnt {
			cnt = 2 * (j + 1)
			res = s[i-j+1 : i+1+j]
		}
	}
	return res
}
