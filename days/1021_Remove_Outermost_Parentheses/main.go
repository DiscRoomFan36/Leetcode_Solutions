package main

import "strings"

func removeOuterParentheses(s string) string {
    counter := 0;
    start := 0;

    sb := strings.Builder{};

    for i, c := range s {
        if c == '(' {
            if counter == 0 {
                start = i;
            }
            counter += 1;
        } else {
            if c != ')' { panic("what?"); }

            counter -= 1;
            if counter == 0 {
                sb.WriteString(s[start+1:i])
            }
        }
    }
    return sb.String();
}
