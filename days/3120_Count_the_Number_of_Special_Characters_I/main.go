package main

import "strings"

func numberOfSpecialChars(word string) int {
    result := 0;
    letters := [26]bool{};

    for _, c := range word {
        if 'A' <= c && c <= 'Z' {
            if letters[c - 'A'] { continue; }
            letters[c - 'A'] = true;

            lower := c - 'A' + 'a';
            if strings.ContainsRune(word, lower) { result += 1; }
        }
    }

    return result;
}
