package main

func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A);
    result := make([]int, n);

    // if item is in A, add 1 to the count, if B subtract 1.
    //
    // we can tell if they have something in common,
    // if the number approaches 0,
    num_to_count := make(map[int]int);
    running_count := 0;

    for i := range n {
        if num_to_count[A[i]] < 0 { running_count += 1; }
        num_to_count[A[i]] += 1;

        if num_to_count[B[i]] > 0 { running_count += 1; }
        num_to_count[B[i]] -= 1;

        result[i] = running_count;
    }

    return result;
}
