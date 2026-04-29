package main

func twoSum(numbers []int, target int) []int {
    left := 0;
    right := len(numbers)-1;

    for left < right {
        for numbers[left] + numbers[right] < target {
            left += 1;
        }

        for numbers[left] + numbers[right] > target {
            right -= 1;
        }

        if numbers[left] + numbers[right] == target {
            return []int{left+1, right+1};
        }

        // since a solution is guarenteed to exist, we dont need to do any bounds checking anywhere.
    }
    panic("not possible.");
}
