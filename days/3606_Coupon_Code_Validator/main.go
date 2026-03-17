package main

import (
	"cmp"
	"slices"
	"unicode"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {

    valid_business_lines_in_order := [4]string{ "electronics", "grocery", "pharmacy", "restaurant" };

    valid_coupon_indexes := make([]int, 0);

    for i := range code {
        is_valid_coupon := true;

        // check active.
        is_valid_coupon = is_valid_coupon && isActive[i];

        if is_valid_coupon {
            // check businessLine
            has_business_line := false;
            for _, business_line := range valid_business_lines_in_order {
                if businessLine[i] == business_line {
                    has_business_line = true;
                    break;
                }
            }
            is_valid_coupon = is_valid_coupon && has_business_line;
        }

        if is_valid_coupon {
            // check code is made of correct chars
            good_code := len(code[i]) > 0;
            for _, c := range code[i] {
                good_char := unicode.IsDigit(c) || unicode.IsLetter(c) || c == '_';
                if !good_char {
                    good_code = false;
                    break;
                }
            }
            is_valid_coupon = is_valid_coupon && good_code;
        }

        if is_valid_coupon {
            Append(&valid_coupon_indexes, i);
        }
    }

    slices.SortFunc(valid_coupon_indexes,
        func(a_index, b_index int) int {
            a_business := businessLine[a_index];
            b_business := businessLine[b_index];
            compare := cmp.Compare(a_business, b_business);
            if compare != 0 { return compare; }

            return cmp.Compare(code[a_index], code[b_index]);
        },
    );

    result := make([]string, len(valid_coupon_indexes));
    for i, coupon_index := range valid_coupon_indexes {
        result[i] = code[coupon_index];
    }
    return result;
}


func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...);
}