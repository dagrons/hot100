package i42

/*
	这里有个很好的题解：https://www.acwing.com/solution/content/34623/
*/

func trap(height []int) int {
	a := height
	ms := []int{}

	res := 0
	for i := 0; i < len(a); i++ {
		for len(ms) > 0 && a[ms[len(ms)-1]] > a[i] {
		}
	}
	return res
}
