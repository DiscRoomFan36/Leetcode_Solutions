package main

func smallestIndex(nums []int) int {
    for i, n := range nums {
        digit_sum := 0;
        for n > 0 && digit_sum <= i {
            digit_sum += n % 10;
            n /= 10;
        }
        if digit_sum == i { return i; }
    }
    return -1;
}
