package main

func minimumDeletions(nums []int) int {
    n := len(nums);
    if n <= 2 { return n; }

    min_index, max_index := 0, 0;
    for i := 1; i < n; i += 1 {
        num := nums[i];

        // numbers in nums are distinct
        if num < nums[min_index] { min_index = i; }
        if num > nums[max_index] { max_index = i; }
    }

    // make min_index always small.
    if max_index < min_index {
        min_index, max_index = max_index, min_index;
    }

    // these are the results after removeing the min and max index,
    //
    // larget array == less we had to remove.
    upto_first := nums[:min_index];
    middle     := nums[min_index+1:max_index];
    to_last    := nums[max_index+1:];

    return n - max(len(upto_first), len(middle), len(to_last));
}
