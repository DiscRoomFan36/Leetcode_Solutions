package main

func kLengthApart(nums []int, k int) bool {
    last_one_position := -(k+1);
    for i, n := range nums {
        if n != 1 { continue; }

        if i - last_one_position <= k { return false; }
        last_one_position = i;
    }
    return true;
}
