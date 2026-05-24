package main

func maximum69Number (num int) int {
    reversed_digits := to_reversed_digits(num);
    for i := len(reversed_digits)-1; i >= 0; i-- {
        if reversed_digits[i] == 6 {
            reversed_digits[i] = 9;
            break;
        }
    }

    result := 0;
    for i := len(reversed_digits)-1; i >= 0; i-- {
        result = result * 10 + reversed_digits[i];
    }
    return result;
}

func to_reversed_digits(x int) []int {
    result := make([]int, 0);
    if x == 0 {
        Append(&result, 0);
    } else {
        for x != 0 {
            Append(&result, x % 10);
            x /= 10;
        }
    }
    return result;
}

func Append[T any](slice *[]T, items ...T) *T {
	*slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}
