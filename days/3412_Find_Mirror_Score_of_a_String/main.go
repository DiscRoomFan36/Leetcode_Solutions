package main

func pop[T ~[]E, E any](slice *T) E {
	item := (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
	return item
}

func calculateScore(s string) int64 {
	// its a map from the letters we've seen, to there index,
	// the last element must have been added most recently
	seen_letters := make([][]int, 26)
	for i := range seen_letters {
		seen_letters[i] = make([]int, 0)
	}

	score := 0

	for i := range s {
		c := s[i] - 'a'
		c_mirror := ('z' - 'a') - c

		if len(seen_letters[c_mirror]) > 0 {
			// pop the last element
			j := pop(&seen_letters[c_mirror])
			score += i - j
		} else {
			seen_letters[c] = append(seen_letters[c], i)
		}
	}

	// this is fine, in go, int is allready 64-bit probably.
	return int64(score)
}
