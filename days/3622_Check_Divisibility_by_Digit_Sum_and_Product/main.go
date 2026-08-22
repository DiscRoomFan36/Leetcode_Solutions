package main

func checkDivisibility(n int) bool {
    inital_n := n;
    digit_sum := 0;
    digit_product := 1;
    for n > 0 {
        digit := n % 10;
        digit_sum     += digit;
        digit_product *= digit;

        n /= 10;
    }

    sum := digit_sum + digit_product;
    if sum == 0 { return false; }

    return inital_n % sum == 0;
}
