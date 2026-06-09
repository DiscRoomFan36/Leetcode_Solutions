package main

func maxTotalValue(nums []int, k int) int64 {
    the_min := nums[0];
    the_max := nums[0];
    for _, n := range nums {
        the_min = min(the_min, n);
        the_max = max(the_max, n);
    }

    return int64((the_max - the_min) * k);
}
