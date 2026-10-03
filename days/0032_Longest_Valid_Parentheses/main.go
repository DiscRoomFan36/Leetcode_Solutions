package main

func longestValidParentheses(s string) int {
    best_len := 0;

    for i := range len(s) {

        // find the longest valid sequence starting from here
        counter := 0;
        for j := i; j < len(s); j += 1 {
            if s[j] == '(' { counter += 1;
            } else {
                if counter == 0 { break; }
                counter -= 1;
            }

            if counter == 0 { best_len = max(best_len, j - i + 1); }
            // haha i chessed it with this line
            if counter > len(s) - j - 1 { break; }
        }
    }

    // the smarter way to do this is to realize that "()" must appear in the longest subsequence,
    // so just find all of those and search outwords

    return best_len;
}
