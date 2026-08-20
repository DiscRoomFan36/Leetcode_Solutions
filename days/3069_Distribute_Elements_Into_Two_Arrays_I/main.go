package main

func resultArray(nums []int) []int {
    array_1 := []int{nums[0]};
    array_2 := []int{nums[1]};

    last := func(array []int) int {
        return array[len(array)-1];
    }

    for i := 2; i < len(nums); i += 1 {
        if last(array_1) > last(array_2) {
            array_1 = append(array_1, nums[i]);
        } else {
            array_2 = append(array_2, nums[i]);
        }
    }

    return append(array_1, array_2...);
}
