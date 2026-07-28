package main

import "strings"

func smallestPalindrome(s string) string {
    n := len(s);
    letter_counts := [26]int{};
    for i := range n/2 {
        letter_counts[s[i] - 'a'] += 1;
    }

    sb := strings.Builder{};
    { // construct
        for i, count := range letter_counts {
            for range count {
                sb.WriteByte(byte(i + 'a'));
            }
        }

        if n % 2 != 0 {
            sb.WriteByte(s[n/2]);
        } 

        for i := len(letter_counts)-1; i >= 0; i-- {
            count := letter_counts[i];
            for range count {
                sb.WriteByte(byte(i + 'a'));
            }
        }
    }
    return sb.String();
}
