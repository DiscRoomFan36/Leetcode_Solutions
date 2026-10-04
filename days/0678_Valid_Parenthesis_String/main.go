package main

func checkValidString(s string) bool {
    memo := make([]int, len(s) * len(s));

    var check_valid_string func(index int, open_count int) bool;
    check_valid_string = func(index int, open_count int) bool {
        if index == len(s) { return open_count == 0; }

        memo_index := open_count * len(s) + index;
        memod := memo[memo_index];
        if memod != 0 { return memod == 1; }

        is_valid := false;
        switch s[index] {
            case '*': {
                is_valid = is_valid || check_valid_string(index + 1, open_count + 1);
                if open_count > 0 { is_valid = is_valid || check_valid_string(index + 1, open_count - 1); }
                is_valid = is_valid || check_valid_string(index + 1, open_count);
            }
            case '(': {
                is_valid = is_valid || check_valid_string(index + 1, open_count + 1);
            }
            case ')': {
                if open_count > 0 {
                    is_valid = is_valid || check_valid_string(index + 1, open_count - 1);
                }
            }
        }

        if is_valid {
            memo[memo_index] = 1;
        } else {
            memo[memo_index] = 2;
        }
        return is_valid;
    }

    return check_valid_string(0, 0);
}
