package main

func totalWaviness(num1 int, num2 int) int {
    total := 0;
    for i := max(num1, 100); i <= num2; i++ {
        num := i;
        digits := []int{};
        for num != 0 {
            digits = append(digits, num % 10);
            num /= 10;
        }

        waviness := 0;
        for j := 1; j < len(digits)-1; j++ {
            a := digits[j-1];
            b := digits[j  ];
            c := digits[j+1];

            if b < a && b < c { waviness += 1; }
            if b > a && b > c { waviness += 1; }
        }

        total += waviness;
    }
    return total;
}
