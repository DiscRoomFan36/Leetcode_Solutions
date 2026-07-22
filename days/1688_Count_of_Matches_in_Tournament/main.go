package main

func numberOfMatches(n int) int {
    matches := 0;
    for n != 1 {
        matches += n / 2;
        // round up
        n = (n + 1) / 2;
    }
    return matches;
}
