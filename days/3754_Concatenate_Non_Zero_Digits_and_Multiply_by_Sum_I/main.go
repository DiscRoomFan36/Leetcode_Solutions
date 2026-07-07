package main

func sumAndMultiply(n int) int64 {
    pow_10 := 1;
    x := 0;
    sum := 0;
    for n > 0 {
        digit := n % 10;
        if digit != 0 {
            x = x + pow_10 * digit;
            pow_10 *= 10;
            sum += digit;
        }

        n /= 10;
    }
    return int64(x * sum);
}
