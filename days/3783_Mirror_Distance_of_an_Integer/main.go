package main

func mirrorDistance(n int) int {
    return Abs(n - reverse(n));
}

func reverse(x int) int {
    result := 0;
    for x > 0 {
        result = result * 10 + x % 10;
        x /= 10;
    }
    return result;
}

func Abs(x int) int {
    if x < 0 { x *= -1; }
    return x;
}
