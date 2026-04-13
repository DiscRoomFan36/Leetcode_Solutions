package main

func getMinDistance(nums []int, target int, start int) int {
    for index := range nums {
        i := start + index;
        j := start - index;

        if i >= len(nums) && j < 0 { break; }

        if j >= 0 {
            if nums[j] == target { return index; }
        }
        if i < len(nums) {
            if nums[i] == target { return index; }
        }
    }

    panic("no target in nums");
}
