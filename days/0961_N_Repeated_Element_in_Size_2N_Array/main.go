package main

func repeatedNTimes(nums []int) int {
    seen_before := make(map[int]bool)
    for _, n := range nums {
        if seen_before[n] { return n }
        seen_before[n] = true
    }
    panic("no repeats found!")
}
