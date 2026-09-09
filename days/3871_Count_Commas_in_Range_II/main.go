package main

func countCommas(_n int64) int64 {
    n := int(_n);
    result := 0;

    pow_1000 := 1_000;
    comma_count := 0;
    for n >= pow_1000 {
        result += (pow_1000 - pow_1000/1000) * comma_count;
        pow_1000 *= 1000;
        comma_count += 1;
    }

    // handle remaining
    result += ((n - pow_1000/1000) + 1) * comma_count;
    return int64(result);
}
