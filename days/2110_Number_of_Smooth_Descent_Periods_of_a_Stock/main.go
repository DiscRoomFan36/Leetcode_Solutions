package main

func getDescentPeriods(prices []int) int64 {
    result := 0

    for i := 0; i < len(prices); {
        // find how long this run is.
        decent_length := 1
        for j := i+1; j < len(prices); j++ {
            if prices[j] != prices[j-1]-1 { break }
            decent_length += 1
        }

        // the proof is something like this:
        //
        // "there are n groups of 1 elements, n-1 groups of 2 elements, and so forth until 1 group of n elements"
        // n + n-1 + ... + 1 -> sum_of_all_digits_upto(n)
        result += sum_of_all_digits_upto(decent_length)

        i += decent_length
    }

    return int64(result)
}

// classic identity of 
// 1 + 2 + 3 + ... + n = (n * (n+1)) / 2
func sum_of_all_digits_upto(n int) int {
    return (n * (n+1)) / 2
}