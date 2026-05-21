package main

func longestCommonPrefix(arr1 []int, arr2 []int) int {
    // make arr1 the smaller array
    if len(arr1) > len(arr2) { arr1, arr2 = arr2, arr1; }

    trie := Digit_Trie{};
    for _, num := range arr1 {
        trie.add_number(num);
    }

    best_prefix := 0;
    for _, num := range arr2 {
        reversed_digits := to_reversed_digits(num);
        prefix := 0;
        current_trie := &trie;
        for i := len(reversed_digits)-1; i >= 0; i-- {
            digit := reversed_digits[i];

            current_trie = current_trie.next[digit];
            if current_trie == nil { break; }
            prefix += 1;
        }

        best_prefix = max(best_prefix, prefix);
    }
    return best_prefix;
}

func to_reversed_digits(x int) []int {
    result := make([]int, 0);
    if x == 0 {
        Append(&result, 0);
    } else {
        for x != 0 {
            Append(&result, x % 10);
            x /= 10;
        }
    }
    return result;
}

type Digit_Trie struct {
    next [10]*Digit_Trie;
}

func (trie *Digit_Trie) add_number(x int) {
    reversed_digits := to_reversed_digits(x);

    current_trie := trie;
    for i := len(reversed_digits)-1; i >= 0; i-- {
        digit := reversed_digits[i];

        slot := &current_trie.next[digit];
        if *slot == nil {
            *slot = new(Digit_Trie);
        }

        current_trie = *slot;
    }
}

func Append[T any](slice *[]T, items ...T) *T {
	*slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}
