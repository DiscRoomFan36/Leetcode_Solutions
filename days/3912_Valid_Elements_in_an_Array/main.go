package main

func findValidElements(nums []int) []int {
    is_valid := make([]bool, len(nums));
    {
        maximum := 0;
        for i := 0; i < len(nums); i++ {
            n := nums[i];
            if n > maximum { is_valid[i] = true; }
            maximum = max(maximum, n);
        }
    }
    {
        maximum := 0;
        for i := len(nums)-1; i >= 0; i-- {
            n := nums[i];
            if n > maximum { is_valid[i] = true; }
            maximum = max(maximum, n);
        }
    }

    result := make([]int, 0);
    for i, n := range nums {
        if is_valid[i] {
            result = append(result, n);
        }
    }
    return result;
}
