package main

type Trie struct {
	next [26]*Trie
}

func minValidStrings(words []string, target string) int {
	const UNSET = 9999999

	root := Trie{}

	// add all words to Trie
	for _, word := range words {
		head := &root
		for _, c := range word {
			if head.next[c-'a'] == nil {
				head.next[c-'a'] = new(Trie)
			}
			head = head.next[c-'a']
		}
	}

	// keeps track of how many steps until you reach the end.
	dp := make([]int, len(target)+1)
	for i := range dp {
		dp[i] = UNSET
	}

	var recur func(words []string, target string) int
	recur = func(words []string, target string) int {
		if dp[len(target)] != UNSET {
			return dp[len(target)]
		}

		result := -1

		best_prefix_length := 0
		head := &root
		for _, char := range target {
			next := head.next[char-'a']
			if next == nil {
				break
			}
			best_prefix_length += 1
			head = next
		}

		if best_prefix_length == len(target) {
			dp[len(target)] = 1
			return 1
		}

		for j := range best_prefix_length {
			to_take := best_prefix_length - j

			recur_res := recur(words, target[to_take:])
			// didn't work out.
			if recur_res == -1 {
				continue
			}

			steps_to_the_end := recur_res + 1

			if (result == -1) || (steps_to_the_end < result) {
				result = steps_to_the_end
			}
		}

		dp[len(target)] = result
		return result
	}

	return recur(words, target)
}
