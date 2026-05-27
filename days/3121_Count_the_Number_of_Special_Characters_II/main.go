package main

func numberOfSpecialChars(word string) int {
    set := make(map[rune]struct{});
    set_of_special := make(map[rune]struct{});

    for _, c := range word {
        if 'A' <= c && c <= 'Z' {
            // second or more occur of upper c
            if Contains(set, c) { continue; }
            set[c] = struct{}{};

            lower := c - 'A' + 'a';
            if Contains(set, lower) {
                // this is the first occurence! and its after all lower! (probably)
                set_of_special[c] = struct{}{};
            }

        } else {
            set[c] = struct{}{};

            // remove the upper if it was there.
            upper := c - 'a' + 'A';
            delete(set_of_special, upper);
        }
    }
    return len(set_of_special);
}

func Contains[T comparable, U any](the_map map[T]U, key T) bool {
    _, found := the_map[key];
    return found;
}
