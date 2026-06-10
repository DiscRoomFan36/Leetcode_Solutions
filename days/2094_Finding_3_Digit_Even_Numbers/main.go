package main

import "slices"

func findEvenNumbers(digits []int) []int {
    all_combos := make(map[int]struct{});

    combine_3 := func(a, b, c int) int {
        return a * 100 + b * 10 + c;
    }

    add_digit_combos := func(a, b, c int) {
        if a > b { a, b = b, a; }
        if b > c { b, c = c, b; }
        if a > b { a, b = b, a; }

        // dont test all combos if we've already checked them
        _, found := all_combos[combine_3(a, b, c)];
        if found { return; }

        // TODO sort a, b, c and dont run again if the same digits appear
        test_combo := func(a, b, c int) {
            if a == 0     { return; }
            if c % 2 != 0 { return; }

            all_combos[combine_3(a, b, c)] = struct{}{};
        }

        test_combo(a, b, c);
        test_combo(a, c, b);
        test_combo(b, a, c);
        test_combo(b, c, a);
        test_combo(c, a, b);
        test_combo(c, b, a);
    }

    // loop over combos in digit array
    for i := 0; i < len(digits); i++ {
        for j := i+1; j < len(digits); j++ {
            for k := j+1; k < len(digits); k++ {
                add_digit_combos(digits[i], digits[j], digits[k]);
            }
        }
    }

    result := make([]int, 0, len(all_combos));
    for k, _ := range all_combos {
        result = append(result, k);
    }
    slices.Sort(result);
    return result;
}
