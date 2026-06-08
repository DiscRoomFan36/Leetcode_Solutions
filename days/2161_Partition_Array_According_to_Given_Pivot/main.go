package main

func pivotArray(nums []int, pivot int) []int {
    less_count  := 0;
    equal_count := 0;
    for _, n := range nums {
        switch true {
            case n <  pivot: { less_count  += 1; }
            case n == pivot: { equal_count += 1; }
        }
    }

    result := make([]int, len(nums));
    less_index  := 0;
    equal_index := less_count;
    more_index  := less_count + equal_count;

    for _, n := range nums {
        switch true {
            case n <  pivot: { result[less_index ] = n; less_index  += 1; }
            case n == pivot: { result[equal_index] = n; equal_index += 1; }
            case n >  pivot: { result[more_index ] = n; more_index  += 1; }
        }
    }

    return result;
}
