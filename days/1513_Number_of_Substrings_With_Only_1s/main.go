package main

func numSub(s string) int {
    const MOD = 1000000007;
    result := 0;
    in_a_row := 0;
    for _, c := range s {
        if c == '1' {
            in_a_row += 1;
            result = (result + in_a_row) % MOD;
        } else {
            in_a_row = 0;
        }
    }
    return result;
}