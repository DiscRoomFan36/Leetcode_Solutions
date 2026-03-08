package main

import "fmt"

func findDifferentBinaryString(nums []string) string {
    n := len(nums);

    all_options := make([]bool, 1 << n);

    for _, num := range nums {
        binary_num := 0;
        for i := range num {
            zero_or_one := int(num[i] - '0');
            binary_num = (binary_num << 1) | zero_or_one;
        }

        all_options[binary_num] = true;
    }

    for i, is_in_nums := range all_options {
        if !is_in_nums {
            // format magic
            return fmt.Sprintf("%0*b", n, i);
        }
    }

    panic("not possible.");
}