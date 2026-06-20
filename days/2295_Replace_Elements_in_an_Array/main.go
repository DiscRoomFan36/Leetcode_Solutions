package main

func arrayChange(nums []int, operations [][]int) []int {
    num_to_index := make(map[int]int);
    for i, n := range nums { num_to_index[n] = i; }

    for _, operation := range operations {
        num, replace_with := operation[0], operation[1];

        // find where the num currently lives.
        index_to_replace := num_to_index[num];
        // replace the number in the array
        nums[index_to_replace] = replace_with;
        // remove old entry
        delete(num_to_index, num);
        // add new entry
        num_to_index[replace_with] = index_to_replace;
    }

    return nums;
}
