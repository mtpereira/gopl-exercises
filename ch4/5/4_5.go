package ch4

func removeAdjacentDuplicates(s []string) []string {
	last := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[last] {
			last++
			s[last] = s[i]
		}
	}
	return s[:last+1]
}
