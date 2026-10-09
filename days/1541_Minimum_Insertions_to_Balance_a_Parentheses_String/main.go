package main

func minInsertions(s string) int {
    // this question sucks ballz

    insertions := 0;
    left_count := 0;
    index      := 0;
    for index < len(s) {
        c := s[index];
        if c == '(' {
            left_count += 1;
            index += 1;
        } else {
            if left_count > 0 {
                left_count -= 1;
            } else {
                insertions += 1;
            }

            if index < len(s)-1 && s[index+1] == ')' {
                index += 2;
            } else {
                insertions += 1;
                index += 1;
            }
        }
    }
    return insertions + left_count * 2;
}
