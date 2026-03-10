package main

func hasIncreasingSubarrays(nums []int, k int) bool {
    for a := range nums {
        b := a + k;
        if b + k > len(nums) { break; }

        if is_increasing(nums[a:a+k]) && is_increasing(nums[b:b+k]) {
            return true;
        }

    }

    return false;
}

func is_increasing(nums []int) bool {
    for i := range len(nums)-1 {
        if nums[i] >= nums[i+1] {
            return false;
        }
    }
    return true;
}