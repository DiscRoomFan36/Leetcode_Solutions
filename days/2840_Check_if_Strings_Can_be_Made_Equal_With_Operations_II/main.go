package main

func checkStrings(s1 string, s2 string) bool {
    s1_even := get_even_digit_letter_freq(s1[0:]);
    s1_odd  := get_even_digit_letter_freq(s1[1:]);

    s2_even := get_even_digit_letter_freq(s2[0:]);
    s2_odd  := get_even_digit_letter_freq(s2[1:]);

    return (s1_even == s2_even) && (s1_odd == s2_odd);
}

func get_even_digit_letter_freq(s string) [26]int {
    result := [26]int{};
    for i := 0; i < len(s); i += 2 {
        c := s[i];
        result[c - 'a'] += 1;
    }
    return result;
}
