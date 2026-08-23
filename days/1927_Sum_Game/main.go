package main

func sumGame(num string) bool {
    get_sum_and_count := func(s string) (int, int) {
        sum, count := 0, 0;
        for _, c := range s {
            if c == '?' {
                count += 1;
            } else {
                sum += int(c - '0');
            }
        }
        return sum, count;
    }

    n0, q0 := get_sum_and_count(num[:len(num)/2]);
    n1, q1 := get_sum_and_count(num[ len(num)/2:]);

    // // if odd, alice has the last laught.
    if (q0 + q1) % 2 != 0 { return true; }

    // no clue why this works, something about
    // bob being able to match alices moves perfectly.
    return (n0 - n1) != (q1 - q0) * 9 / 2;
}
