package main

import "math"

func minElement(nums []int) int {
    min_element := math.MaxInt;
    for _, n := range nums {
        min_element = min(min_element, sum_digits(n));
    }
    return min_element;
}

func sum_digits(n int) int {
    digit_sum := 0;
    for n > 0 {
        digit_sum += n % 10;
        n /= 10;
    }
    return digit_sum;
}
