package main

func isGood(nums []int) bool {
    n := len(nums)-1;

    if n == 0 { return false; }
    in_array := make([]bool, n-1);

    n_count := 0;
    for _, num := range nums {
        if num == n {
            n_count += 1;
            continue;
        }

        if num <= 0 || num > n { return false; }

        if in_array[num-1] { return false; }
        in_array[num-1] = true;
    }

    if n_count != 2 { return false; }

    return true;
}
