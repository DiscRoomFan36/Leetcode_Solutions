package main

func largestAltitude(gain []int) int {
    running_sum := 0;
    best := 0;
    for _, n := range gain {
        running_sum += n;
        best = max(best, running_sum);
    }
    return best;
}
