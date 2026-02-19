package main

import "strings"

// propbably overcomplicated this.
func countBinarySubstrings(s string) int {
    if len(s) <= 1 { return 0; }

    first_char := s[0];
    other_char := byte('0');
    if first_char == '0' { other_char = '1'; }

    // find the first different char,
    first_difference := strings.IndexByte(s, other_char);
    if first_difference == -1 { return 0; }

    result := 0;

    for k := first_difference-1; k < len(s)-1; {
        c1, c2 := s[k], s[k+1];

        // only look for the edges, then backtrack until we hit a different thing
        if c1 == c2 { panic("not possible"); }

        result += 1;

        // search outwards
        i, j := k-1, k+2;
        for (i >= 0) && (j < len(s)) && (s[i] == c1) && (s[j] == c2) {
            result += 1;
            i -= 1;
            j += 1;
        }

        // skip all the same characters.
        next_difference := strings.IndexByte(s[j:], c1);
        if next_difference == -1 { break; }
        k = j + next_difference - 1;
    }

    return result;
}