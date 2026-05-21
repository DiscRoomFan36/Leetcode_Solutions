package main

func getMaximumGenerated(n int) int {
    if n == 0 { return 0; }

    nums := make([]int, n+1);

    nums[0] = 0;
    nums[1] = 1;

    for i := 2; i <= n; i++ {
        if i % 2 == 0 {
            nums[i] = nums[i/2];
        } else {
            nums[i] = nums[i / 2] + nums[i / 2 + 1];
        }
    }

    maximum := 1;
    for _, num := range nums {
        maximum = max(maximum, num); 
    }
    return maximum;
}
