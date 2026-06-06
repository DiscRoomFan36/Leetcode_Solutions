package main

func leftRightDifference(nums []int) []int {
    result := make([]int, len(nums));

    for i := range nums {
        left_sum := 0;
        for _, n := range nums[:i] { left_sum  += n; }
        right_sum := 0;
        for _, n := range nums[i+1:] { right_sum += n; }

        result[i] = Abs(left_sum - right_sum);
    }

    return result;
}

func Abs(x int) int {
    if x < 0 { x *= -1; }
    return x;
}
