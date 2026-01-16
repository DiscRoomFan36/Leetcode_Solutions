package main

func countGoodSubstrings(s string) int {
	result := 0
	for i := range len(s) - 2 {
		a, b, c := s[i], s[i+1], s[i+2]
		if a != b && a != c && b != c {
			result += 1
		}
	}
	return result
}
