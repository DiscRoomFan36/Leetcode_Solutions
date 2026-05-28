package main

func firstUniqChar(s string) int {
    type Foo struct {
        index int;
        count int;
    };

    characters := make(map[rune]Foo);

    for i, c := range s {
        foo := characters[c];
        foo.index = i;
        foo.count += 1;
        characters[c] = foo;
    }

    best_index := -1;
    for _, foo := range characters {
        if foo.count == 1 {
            if best_index == -1 || foo.index < best_index {
                best_index = foo.index;
            }
        }
    }
    return best_index;
}
