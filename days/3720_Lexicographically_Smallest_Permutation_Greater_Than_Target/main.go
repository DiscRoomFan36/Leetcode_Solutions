package main

func lexGreaterPermutation(s string, target string) string {
    var recur func(current_counts [26]int, index int, is_over_target bool, output []byte) bool;
    recur = func(current_counts [26]int, index int, is_over_target bool, output []byte) bool {
        if index == len(target) {
            // strictly greater
            return is_over_target;
        }

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
    }

    input_letter_count := get_letter_counts(s);
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
