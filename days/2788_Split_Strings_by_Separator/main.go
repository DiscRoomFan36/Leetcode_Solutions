package main

import "strings"

func splitWordsBySeparator(words []string, separator byte) []string {
    result := make([]string, 0);
    for _, word := range words {
        split_word := strings.Split(word, string(separator));
        for _, part := range split_word {
            if part != "" {
                result = append(result, part);
            }
        }
    }
    return result;
}
