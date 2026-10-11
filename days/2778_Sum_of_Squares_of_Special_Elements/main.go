package main

func sumOfSquares(nums []int) int {
    sum := 0;
    for i, n := range nums {
        if len(nums) % (i+1) == 0 {
            sum += n*n;
        }
    }
    return sum;
}
