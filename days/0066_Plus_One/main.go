package main

func plusOne(digits []int) []int {
    for i := len(digits)-1; i >= 0; i-- {
        digits[i] = (digits[i] + 1) % 10;
        // check if we need to carry
        if digits[i] != 0 {
            return digits;
        }
    }

    // we need to put a 1 at the front.
    return append([]int{1}, digits...);
}