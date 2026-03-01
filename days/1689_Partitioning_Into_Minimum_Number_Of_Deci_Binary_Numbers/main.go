package main

func minPartitions(n string) int {
    max_digit := 0;
    for _, c := range n {
        digit := int(c - '0');
        max_digit = max(max_digit, digit);
        if max_digit == 9 { break; }
    }
    return max_digit;
}
