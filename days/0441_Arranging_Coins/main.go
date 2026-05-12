package main

func arrangeCoins(n int) int {
    step := 0;
    for n >= 0 {
        step += 1;
        n -= step;
    }
    return step-1;
}
