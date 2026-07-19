package main

import "strings"

func removeDuplicates(s string, k int) string {
    type Foo struct {
        letter byte;
        count int;
    }

    letters_and_counts := []Foo{};
    for i := range s {
        c := s[i];
        if len(letters_and_counts) == 0 {
            Append(&letters_and_counts, Foo{c, 1});
            continue;
        }

        last_index := len(letters_and_counts)-1;
        last := &letters_and_counts[last_index];
        if last.letter == c {
            last.count += 1;
            if last.count >= k {
                // pop
                letters_and_counts = letters_and_counts[:last_index];
            }
        } else {
            Append(&letters_and_counts, Foo{c, 1});
        }
    }

    sb := strings.Builder{};
    for _, it := range letters_and_counts {
        for range it.count { sb.WriteByte(it.letter); }
    }
    return sb.String();
}

func Append[T any](slice *[]T, items ...T) *T {
	*slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}
