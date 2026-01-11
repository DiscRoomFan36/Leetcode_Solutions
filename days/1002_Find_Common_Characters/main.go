package main

import "fmt"

const LETTERS_IN_ALPHABET = 'z' - 'a' + 1

// uint8 because "1 <= words[i].length <= 100"
type Letter_Count_Array [LETTERS_IN_ALPHABET]uint8

func word_to_letter_counts(word string) Letter_Count_Array {
	result := Letter_Count_Array{}
	for _, c := range word {
		if !('a' <= c && c <= 'z') {
			panic(fmt.Errorf("letter is not a lowercase english letter, letter: %b, word: %v", c, word))
		}
		result[c-'a'] += 1
	}
	return result
}

func commonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}
	common_chars := word_to_letter_counts(words[0])
	for _, word := range words[1:] {
		counts := word_to_letter_counts(word)

		for i := range common_chars {
			// this should be efficient. there are assembly instructions for this exact case.
			common_chars[i] = min(common_chars[i], counts[i])
		}
	}

	result := make([]string, 0)
	for i := range common_chars {
		for range common_chars[i] {
			// don't like this Sprintf, but its what the question requires
			result = append(result, fmt.Sprintf("%c", i+'a'))
		}
	}
	return result
}
