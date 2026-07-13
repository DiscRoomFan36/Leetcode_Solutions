package main

func sequentialDigits(low int, high int) []int {
    log_10 := int_log_10(low);
    pow_10 := int_pow_10(log_10);

    result := make([]int, 0);

    outer: for {
        // number > high test could be missed.
        if pow_10 > high { break; }

        inner: for i := 1; i <= 9; i++ {
            // digits wont increase, they wrap to zero.
            if i + log_10 > 9 { break inner; }
            number := 0;
            this_pow_10 := pow_10;
            for j := 0; j <= log_10; j++ {
                number += (i+j) * this_pow_10;
                this_pow_10 /= 10;
            }

            if number < low  { continue inner; }
            if number > high { break    outer; }

            result = append(result, number);
        }

        log_10 += 1;
        pow_10 *= 10;
    }

    return result;
}

func int_pow_10(n int) int {
    return int_pow(10, n);
}

func int_log_10(x int) int {
    if (x <= 0) { return 0; }
    result := 0;
    for x > 0 { result += 1; x /= 10; }
    return result - 1;
}

// dose the fast thing.
func int_pow(x, y int) int {
    if x == 0 { return 0; }
    if y <  0 { return 0; }

    result := 1;
    keep_squaring := x;
    for y != 0 {
        if y & 1 != 0 { result *= keep_squaring; }
        y = y >> 1;
        keep_squaring = keep_squaring*keep_squaring;
    }

    return result;
}
