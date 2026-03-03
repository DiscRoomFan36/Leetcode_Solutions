package main

func findKthBit(n int, k int) byte {
    binary_string := make([]bool, 1, k);

    // for range n-1 {
    // the start of the string dose not change, so only do it until we have enough for k
    for len(binary_string) <= k {
        binary_string = append(binary_string, true);
        for i := len(binary_string)-2; i >= 0; i-- {
            binary_string = append(binary_string, !binary_string[i]);
        }
    }

    if binary_string[k-1] {
        return '1';
    } else {
        return '0';
    }
}