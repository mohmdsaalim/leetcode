func kWeakestRows(mat [][]int, k int) []int {
	m, a := map[int]int{}, []int{}
	for i := 0; i < len(mat); i++ {
		c := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				c++
			}
		}
		m[i] = c
		a = append(a, i)
	}
	sort.Slice(a, func(i, j int) bool {
		if m[a[i]] == m[a[j]] {
			return a[i] < a[j]
		}
		return m[a[i]] < m[a[j]]
	})
	return a[:k]
}