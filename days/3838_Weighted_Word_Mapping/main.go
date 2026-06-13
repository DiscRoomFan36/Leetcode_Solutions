package main

import "strings"

func mapWordWeights(words []string, weights []int) string {
    word_weight := func(word string) int {
        total := 0;
        for _, c := range word {
            total += weights[c - 'a'];
        }
        return total;
    }

    sb := strings.Builder{};
    for _, word := range words {
        weight := word_weight(word) % 26;
        letter := ((26-1) - weight) + 'a';
        sb.WriteRune(rune(letter));
    }
    return sb.String();
}
