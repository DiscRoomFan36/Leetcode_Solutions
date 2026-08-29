package main

import "slices"

func lexicographicallySmallestArray(nums []int, limit int) []int {
    // assuming:
    //    - A and B can swap
    //    - B and C can swap
    //
    // it follows that:
    // A and C can swap.

    // 1. find the groups of things that can swap,
    // 2. sort the numbers in the groups,
    // 3. replace in original array

    n := len(nums);

    copy_of_nums := make([]int, n);
    copy(copy_of_nums, nums);
    slices.Sort(copy_of_nums);

    current_group_number := 0;
    index_to_group := make([]int, n);

    groups := []int{0};
    for i := 1; i < n; i += 1 {
        if Abs(copy_of_nums[i-1] - copy_of_nums[i]) <= limit {
            groups[len(groups)-1] += 1;
        } else {
            groups = append(groups, i);

            current_group_number += 1;
        }
        index_to_group[i] = current_group_number;
    }

    group_start := 0;
    slices_of_groups := make([][]int, 0);

    for _, group_end := range groups {
        slices_of_groups = append(slices_of_groups, copy_of_nums[group_start:group_end+1]);
        group_start = group_end + 1;
    }

    // construct final result
    for i, n := range nums {
        index, found := slices.BinarySearch(copy_of_nums, n);
        if !found { panic("not possible!"); }

        group_index := index_to_group[index];
        slice_of_group := &slices_of_groups[group_index];

        // pop front.
        nums[i] = (*slice_of_group)[0];
        *slice_of_group = (*slice_of_group)[1:];
    }

    return nums;
}

func Abs(x int) int {
    if x < 0 { x = -x; }
    return x;
}
