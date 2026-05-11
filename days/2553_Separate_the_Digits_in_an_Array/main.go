package main

func separateDigits(nums []int) []int {
    result := make([]int, 0);

    for _, num := range nums {
        digits_reversed := make([]int, 0, 32);
        for num > 0 {
            digit := num % 10;
            num /= 10;

            Append(&digits_reversed, digit);
        }

        for i := len(digits_reversed)-1; i >= 0; i-- {
            Append(&result, digits_reversed[i]);
        }
    }

    return result;
}

func Append[T any](slice *[]T, items ...T) *T {
	*slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}
