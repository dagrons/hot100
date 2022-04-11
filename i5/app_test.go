package i5_test

import (
	"testing"

	"github.com/dagrons/hot100/i5"
)

func TestLongestPalindrome(t *testing.T) {
	a := "cbbd"
	rs := i5.LongestPalindrome(a)
	if rs != "bb" {
		t.Errorf("longestPalindrome failed, s = %s, get %s", a, rs)
	}
}
