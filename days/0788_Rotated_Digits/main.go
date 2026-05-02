package main

func rotatedDigits(n int) int {
    num_good := 0;
    for i := 1; i <= n; i++ {
        if good_after_rotaion(i) {
            num_good += 1;
        }
    }
    return num_good;
}

func good_after_rotaion(x int) bool {
    will_be_different := false;

    for x != 0 {
        digit := x % 10;
        x /= 10;

        switch digit {
            case 0, 1, 8: { /* do nothing */ }
            case 2, 5, 6, 9: { will_be_different = true; }
            default: { return false; } // other digets cannot be rotated.
        }
    }

    return will_be_different;
}
