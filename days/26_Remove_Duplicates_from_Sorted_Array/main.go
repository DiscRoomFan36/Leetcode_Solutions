
package main

func removeDuplicates(nums []int) int {
    index := 0
    last_seen_number := -99999 // a number that cannot appear in nums 
    for _, n := range nums {
        // ignore the duplicates
        if n == last_seen_number { continue }

        last_seen_number = n
        nums[index] = n
        index += 1
    }

    return index
}
