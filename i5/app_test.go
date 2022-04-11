package i5_test

import (
	"testing"

	"github.com/dagrons/hot100/i5"
)

func TestLongestPalindrome(t *testing.T) {
	tries := map[string]string{
		"cbbd":  "bb",
		"babad": "bab",
	}
	for k, v := range tries {
		rs := i5.LongestPalindrome(k)
		if rs != v {
			t.Errorf("in=%s, out=%s", k, v)
		}
	}
}
