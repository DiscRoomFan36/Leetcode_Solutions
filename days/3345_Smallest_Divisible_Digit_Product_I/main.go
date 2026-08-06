package main

func smallestNumber(n int, t int) int {
    // 0 is always divisible
    for i := n; ; i += 1 {
        if multiply_digits(i) % t == 0 { return i; }
    }
}

func multiply_digits(n int) int {
    total := 1;
    for n > 0 {
        total *= n % 10;
        n /= 10;
    }
    return total;
}
