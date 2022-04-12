package i32_test

import (
	"testing"

	"github.com/dagrons/hot100/i32"
)

func TestLongestValidParentheses(t *testing.T) {
	tries := map[string]int{
		"(()":    2,
		")()())": 4,
	}
	for k, v := range tries {
		ans := i32.LongestValidParentheses(k)
		if ans != v {
			t.Errorf("in = %s, out=%d", k, v)
		}
	}
}
