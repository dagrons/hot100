package i31_test

import (
	"testing"

	"github.com/dagrons/hot100/i31"
)

func TestNextPermutation(t *testing.T) {
	queies := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	ans := [][]int{
		{1, 3, 2},
		{1, 2, 3},
	}
	for i := range queies {
		i31.NextPermutation(queies[i])
		for j := range queies[i] {
			if queies[i][j] != ans[i][j] {
				t.Errorf("in=%v, out=%v", queies[i], ans[i])
				break
			}
		}
	}
}
