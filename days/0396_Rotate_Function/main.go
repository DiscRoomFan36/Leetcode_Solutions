package main

func maxRotateFunction(nums []int) int {
    n := len(nums);
    if n == 0 { return 0; }

    sum := 0;
    for _, n := range nums { sum += n; }

    F_0 := 0;
    for i := range n { F_0 += i * nums[i]; }

    best := F_0;
    F_x := F_0;
    for i := range n {
        F_x = F_x + sum - n * nums[n-i-1];
        best = max(best, F_x);
    }

    return best;
}
