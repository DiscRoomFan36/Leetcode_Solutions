package main

import "strings"

func addBinary(a string, b string) string {
    if a == "0" { return b }
    if b == "0" { return a }

    sb := strings.Builder{}

    carry := 0
    for i := range max(len(a), len(b)) {
        a_digit := 0
        if (i < len(a)) && (a[len(a)-i-1] == '1') { a_digit = 1 }

        b_digit := 0
        if (i < len(b)) && (b[len(b)-i-1] == '1') { b_digit = 1 }

        result := a_digit + b_digit + carry

        byte_to_write := byte('0')
        if result % 2 == 1 { byte_to_write = '1' }

        sb.WriteByte(byte_to_write)
        carry = (result >> 1)
    }

    if carry != 0 { sb.WriteByte('1') }

    // this is the wrong way around, gotta reverse it
    as_string := []byte(sb.String())
    for i := range len(as_string)/2 {
        j := len(as_string)-i-1
        as_string[i], as_string[j] = as_string[j], as_string[i]
    }

    return string(as_string)
}