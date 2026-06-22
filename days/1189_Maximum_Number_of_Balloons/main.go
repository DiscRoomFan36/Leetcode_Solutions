package main

import "math"

func maxNumberOfBalloons(text string) int {
    ballon_letter_count := letter_counts_of_text("balloon")
    text_letter_count   := letter_counts_of_text(text);

    max_number := math.MaxInt;
    for i, count := range ballon_letter_count {
        if count == 0 { continue; }

        other_count := text_letter_count[i];
        max_number = min(max_number, other_count / count);
    }
    return max_number;
}

func letter_counts_of_text(text string) [26]int {
    letter_counts := [26]int{};
    for _, c := range text {
        letter_counts[c - 'a'] += 1;
    }
    return letter_counts;
}
