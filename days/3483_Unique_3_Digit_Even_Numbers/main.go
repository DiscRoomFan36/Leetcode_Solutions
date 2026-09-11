package main

func totalNumbers(digits []int) int {
    digit_counts := [10]int{};
    for _, digit := range digits {
        digit_counts[digit] += 1;
    }

    result := 0;
    for i := 1; i < 10; i += 1 {
        if digit_counts[i] == 0 { continue; }
        digit_counts[i] -= 1;
        for j := 0; j < 10; j += 1 {
            if digit_counts[j] == 0 { continue; }
            digit_counts[j] -= 1;
            for k := 0; k < 10; k += 2 {
                if digit_counts[k] == 0 { continue; }
                result += 1;
            }
            digit_counts[j] += 1;
        }
        digit_counts[i] += 1;
    }

    return result;
}
