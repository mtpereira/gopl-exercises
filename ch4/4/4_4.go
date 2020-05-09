package ch4

// rotate a slice to the right.
func rotate(s []int, n int) []int {
	if s == nil {
		return nil
	}

	r := make([]int, len(s)+len(s)-n, len(s)+len(s)-n)
	copy(r[:len(s)], s)
	copy(r[len(s):], s[:len(s)-n])
	return r[len(s)-n:]
}
