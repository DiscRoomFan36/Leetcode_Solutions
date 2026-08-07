package main

import "strings"

var primes = [4]int{2, 3, 5, 7};
type Prime_Count [4]int;

func smallestNumber(num string, t_ int64) string {
    t := int(t_);

    // if t has a prime factor > 9, return -1,
    t_prime_counts := Prime_Count{};
    for i, p := range primes {
        for t % p == 0 {
            t_prime_counts[i] += 1;
            t /= p;
        }
    }
    // there must be a prime bigger than 9, we cannot make this.
    if t != 1 { return "-1"; }

    digit_to_prime_factors := []Prime_Count{
        /* 0 */ {0, 0, 0, 0},
        /* 1 */ {0, 0, 0, 0},
        /* 2 */ {1, 0, 0, 0},
        /* 3 */ {0, 1, 0, 0},
        /* 4 */ {2, 0, 0, 0},
        /* 5 */ {0, 0, 1, 0},
        /* 6 */ {1, 1, 0, 0},
        /* 7 */ {0, 0, 0, 1},
        /* 8 */ {3, 0, 0, 0},
        /* 9 */ {0, 2, 0, 0},
    };

    number := make([]byte, len(num));
    make_all_1s := false;
    for i := range len(num) {
        digit := num[i] - '0';
        if digit == 0 {
            // answer cannot have 0's
            // all digits from on are 1's
            make_all_1s = true;
        }
        if make_all_1s {
            number[i] = 1;
        } else {
            number[i] = digit;
        }
    }

    need_at_least_digits := min_digits_needed_to_make(t_prime_counts);
    if need_at_least_digits > len(number) {
        // turn number into smallest number with N digits, aka all 1's
        for i := range number {
            number[i] = 1;
        }
        for len(number) < need_at_least_digits {
            number = append(number, 1);
        }
    }

    // returns "are we done?"
    var recur func(needed_prime_count Prime_Count, number []byte, needs_to_be_bigger bool) bool;
    recur = func(needed_prime_count Prime_Count, number []byte, needs_to_be_bigger bool) bool {
        min_digits := min_digits_needed_to_make(needed_prime_count);
        if min_digits > len(number) { return false; }

        // base case we made it 
        if len(number) == 0 {
            if min_digits != 0 { panic("what"); }
            return true;
        }

        start_index := number[0];
        if !needs_to_be_bigger { start_index = 1; }
        for i := start_index; i <= 9; i += 1 {
            next_needs_to_be_bigger := needs_to_be_bigger && (i == number[0]);
            // change this number and continue 
            new_prime_counts := subtract_prime_counts(needed_prime_count, digit_to_prime_factors[i]);
            got_answer := recur(new_prime_counts, number[1:], next_needs_to_be_bigger);
            if got_answer {
                number[0] = i;
                return true;
            }
        }

        // didnt make it.
        return false;
    }

    for {
        made_a_number := recur(t_prime_counts, number, true);
        if made_a_number { break; }
        // else need more digits because starting number was too big.

        // make into smallest number with more digits
        for i := range number { number[i] = 1; }
        number = append(number, 1);
    }

    // turn number into a number
    sb := strings.Builder{};
    for _, c := range number {
        sb.WriteByte(c + '0');
    }
    return sb.String();
}


func min_digits_needed_to_make(prime_count Prime_Count) int {
    make_positive := func(x *int) { *x = max(*x, 0); }
    two, three, five, seven := prime_count[0], prime_count[1], prime_count[2], prime_count[3];
    make_positive(&two); make_positive(&three); make_positive(&five); make_positive(&seven);
    // must use 1 digit per 5 / 7
    result := five + seven;

    // at best, we can use 8 and 9 to grab lots of 2's and 3's
    result += two / 3 + three / 2;

    // deal with the remainder
    two %= 3;
    three %= 2;

    if two == 0 && three == 0 {
        // no extra digits
        result += 0;
    } else if two == 0 && three == 1 {
        // use a 3/6/9
        result += 1;
    } else if two == 1 && three == 0 {
        // use 2 or 4 or 8,
        result += 1;
    } else if two == 1 && three == 1 {
        // use 6
        result += 1;
    } else if two == 2 && three == 0 {
        // use 4 or 8
        result += 1;
    } else if two == 2 && three == 1 {
        // must use 2 digits, one could be 6
        result += 2;
    } else {
        panic("unreachable");
    }

    return result;
}

func subtract_prime_counts(a, b Prime_Count) Prime_Count {
    for i := range a {
        // no negative allowed
        a[i] = max(a[i] - b[i], 0);
    }
    return a;
}
