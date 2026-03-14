package main

import "strings"

func getHappyString(n int, k int) string {
    result := make([]int, 0, n);
    k_after_recur, result_if_any := recur(result, n, k);

    if k_after_recur > 0 { return ""; }

    if len(result_if_any) != n { panic("not possible"); }

    sb := strings.Builder{}
    for _, c := range result_if_any {
        sb.WriteRune(rune(c) + 'a')
    }
    return sb.String();
}

func recur(result []int, n, k int) (new_k int, result_if_any []int) {
    if n == 0 {
        // found a new combination. dec k, check if this is the one we want.
        k -= 1;
        if k == 0 {
            return k, result;
        } else {
            return k, []int{};
        }
    }

    for i := range 3 {
        // no dup's
        if len(result) > 0 && result[len(result)-1] == i { continue; }

        new_result := append(result, i);

        new_k, result_if_any := recur(new_result, n-1, k);
        if len(result_if_any) != 0 { return 0, result_if_any }

        k = new_k;
    }

    return k, []int{};
}
