package search

func BinarySearch(items []int, n int) int {
	var (
		l, h, m = 0, len(items)-1, 0
	)
	for l <= h {
		m = (h + l) / 2
		if items[m] == n {
			return m
		} else if items[m] < n {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return -1
}