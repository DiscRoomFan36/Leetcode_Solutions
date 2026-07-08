package main

func findMinFibonacciNumbers(k int) int {
    fib_numbers := []int{};
    { // fib numbers grow fast, we dont need that many terms.
        a, b := 1, 1;
        for a <= k {
            c := a + b;
            a, b = c, a;
            fib_numbers = append(fib_numbers, b);
        }
    }

    // greedy method
    count := 0;
    for i := len(fib_numbers)-1; i >= 0; i-- {
        fib := fib_numbers[i];
        if k / fib > 0 {
            count += k / fib;
            k = k % fib;

            if k == 0 { break; }
        }
    }

    return count;
}
