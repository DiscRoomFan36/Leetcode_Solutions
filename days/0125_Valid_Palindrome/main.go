package main

import "unicode"

func isPalindrome(s string) bool {
    n := len(s)

    i, j := 0, n-1

    for i < j {
        // find the next letter / number
        for ; i <  n && !is_alpha(s[i]); i++ {}
        for ; j >= 0 && !is_alpha(s[j]); j-- {}

        if i >= n || j < 0 { break }

        c1 := unicode.ToLower( rune(s[i]) )
        c2 := unicode.ToLower( rune(s[j]) )

        if c1 != c2 { return false }

        i += 1
        j -= 1
    }

    return true
}

func is_alpha(b byte) bool {
    r := rune(b)
    return unicode.IsLetter(r) || unicode.IsNumber(r)
}