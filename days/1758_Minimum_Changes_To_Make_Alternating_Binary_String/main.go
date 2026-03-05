package main

func minOperations(s string) int {
    count := 0;
    for i, c := range s {
        // xor
        if (c == '0') != (i % 2 == 0) {
            count += 1;
        }
    }
    // either we flip them to make 01010... or we flip to make 10101...
    return min(count, len(s) - count);
}
