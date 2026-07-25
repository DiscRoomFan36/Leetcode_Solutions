package main

func maxProduct(n int) int {
    biggest_digit  := 0;
    second_biggest := 0;

    for n != 0 {
        digit := n % 10;
        n /= 10;

        if digit > second_biggest {
            second_biggest = digit;
            // swap if its the new biggest.
            if second_biggest > biggest_digit {
                second_biggest, biggest_digit = biggest_digit, second_biggest;
            }
        }
    }

    return biggest_digit * second_biggest;
}
