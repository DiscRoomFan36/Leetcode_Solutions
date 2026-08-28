package main

func lexPalindromicPermutation(s string, target string) string {
    input_letter_count := get_letter_counts(s);

    seen_odd_number := false;
    seen_odd_number_character := byte(0);
    for i, count := range input_letter_count {
        if count % 2 != 0 {
            // even length cannot have odd letter counts
            if len(s) % 2 == 0 { return ""; }

            // cant be 2 odd numbers
            if seen_odd_number { return ""; }
            seen_odd_number           = true;
            seen_odd_number_character = byte(i + 'a');
        }

        input_letter_count[i] /= 2;
    }

    var recur func(current_counts [26]int, index int, is_over_target bool, output []byte) bool;
    recur = func(current_counts [26]int, index int, is_over_target bool, output []byte) bool {
        if index == len(target) {
            // strictly greater
            return is_over_target;
        }

        // fmt.Println("index:", index, "is_over_target:", is_over_target, "output:", output, "current_counts:", current_counts);

        // palindrom check
        if index < len(target)/2 {
            if !is_over_target {
                target_letter := target[index];
                if current_counts[target_letter - 'a'] > 0 {
                    current_counts[target_letter - 'a'] -= 1;

                    output[index] = target_letter;
                    result := recur(current_counts, index + 1, false, output);
                    if result { return true; }

                    current_counts[target_letter - 'a'] += 1;
                }

                // else find a bigger letter
                for i := target_letter - 'a' + 1; i < 26; i += 1 {
                    if current_counts[i] > 0 {
                        current_counts[i] -= 1;

                        output[index] = i + 'a';
                        result := recur(current_counts, index + 1, true, output);
                        if result != true {
                            panic("what>");
                        }
                        return true;
                    }
                }
                // no bigger letter, this is a failure
                return false;

            } else {

                // were allready over, now do smallest
                for i := range 26 {
                    if current_counts[i] > 0 {
                        current_counts[i] -= 1;
                        output[index] = byte(i + 'a');

                        result := recur(current_counts, index + 1, true, output);
                        if !result {
                            panic("what?");
                        }
                        return true;
                    }
                }
                // there must be enough letters.
                panic("unreachable");
            }

        } else {
            // deal with panendrome.
            char_it_has_to_be := output[len(target)-index-1];
            if len(target)%2 != 0 && index == len(target)/2 {
                char_it_has_to_be = seen_odd_number_character;
            }


            output[index] = char_it_has_to_be;
            if !is_over_target {
                if char_it_has_to_be < target[index] { return false; }

                return recur(current_counts, index + 1, char_it_has_to_be > target[index], output);
            } else {
                return recur(current_counts, index + 1, true, output);
            }
        }

    }

    output := make([]byte, len(target));
    result := recur(input_letter_count, 0, false, output);
    if !result { return ""; }

    return string(output);
}

func get_letter_counts(word string) [26]int {
    result := [26]int{};
    for _, c := range word {
        result[c - 'a'] += 1;
    }
    return result;
}
