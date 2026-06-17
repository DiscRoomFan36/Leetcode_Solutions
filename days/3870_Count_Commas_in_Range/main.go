package main

func countCommas(n int) int {
    if n < 1000 { return 0; }
    return n - 1000 + 1;
}
