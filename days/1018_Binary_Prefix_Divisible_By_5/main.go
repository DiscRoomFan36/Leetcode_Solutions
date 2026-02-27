package main

func prefixesDivBy5(nums []int) []bool {
    result := make([]bool, len(nums));

    rolling_number := 0;
    for i := range nums {
        rolling_number = ((rolling_number << 1) | nums[i]) % 5;
        result[i] = (rolling_number % 5 == 0);
    }

    return result;
}