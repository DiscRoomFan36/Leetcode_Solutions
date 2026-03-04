package main

func hasSameDigits(s string) bool {
    if len(s) < 2 { panic("leetcode"); }

    number_array := make([]int, len(s));
    for i, c := range s { number_array[i] = int(c - '0'); }

    const MOD = 10;
    for len(number_array) > 2 {
        for i := range len(number_array)-1 {
            number_array[i] = (number_array[i] + number_array[i+1]) % MOD;
        }
        // remove the last element.
        number_array = number_array[:len(number_array)-1];
    }

    return number_array[0] % MOD == number_array[1] % MOD;
}