package i5

func LongestPalindrome(s string) string {
	var res string
	n := len(s)
	cnt := 0

	for i := range s {
		j := 0
		for ; i-j >= 0 && i+j < n && s[i-j] == s[i+j]; j++ {
		}
		if j > cnt {
			cnt = j + 1
			res = s[i-j+1 : i+j]
		}
	}

	for i := 0; i < n-1; i++ {
		j := 0
		for ; i-j >= 0 && i+j+1 < n && s[i-j] == s[i+j+1]; j++ {
		}
		if j+1 > cnt {
			cnt = j + 1
			res = s[i-j : i+j+1]
		}
	}
	return res
}
