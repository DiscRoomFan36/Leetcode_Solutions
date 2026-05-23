package main

func check(nums []int) bool {
    // find the maybe smallest element.
    maybe_smallest_index := -1;
    for i := range len(nums)-1 {
        if nums[i] > nums[i+1] {
            maybe_smallest_index = i+1;
            break;
        }
    }

    // array is sorted, and not rotated
    if maybe_smallest_index == -1 { return true; }

    // check if the (now unrotated) array is sorted.
    for i := range len(nums)-1 {
        index_1 := (i +     maybe_smallest_index) % len(nums);
        index_2 := (i + 1 + maybe_smallest_index) % len(nums);

        if nums[index_1] > nums[index_2] { return false; }
    }
    return true;
}
