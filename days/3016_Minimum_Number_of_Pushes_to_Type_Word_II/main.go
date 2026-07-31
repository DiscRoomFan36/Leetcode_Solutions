package main

import "slices"

func minimumPushes(word string) int {
    letter_counts := [26]int{};
    for i := range len(word) {
        letter_counts[word[i] - 'a'] += 1;
    }

    slices.Sort(letter_counts[:]);
    // fmt.Println(letter_counts);

    const NUM_KEYS = 8;

    total_presses := 0;
    for i := len(letter_counts)-1; i >= 0; i-- {
        j := len(letter_counts)-1 - i;

        count := letter_counts[i];
        factor := (j / NUM_KEYS) + 1;

        total_presses += count * factor;
    }
    return total_presses;
}
