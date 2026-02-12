package main

func longestBalanced(s string) int {
    if len(s) == 0 { return 0 }

    best := 1 // smallest substring always works.
    for i := range len(s) {
        type Bucket [26]int

        bucket := Bucket{}
        for j := i; j < len(s); j++ {
            c := s[j] - 'a'
            bucket[c] += 1

            num_letters_in_bucket := j - i + 1
            // we could also keep a count of unique letters, and % num_letters for fast check
            if num_letters_in_bucket > best {
                // check if its all good.
                flag := true
                all_counts := bucket[c] // check if all counts are the same
                for _, letter_count := range bucket {
                    if letter_count == 0 { continue }
                    if letter_count != all_counts {
                        flag = false
                        break
                    }
                }

                if flag { best = num_letters_in_bucket }
            }
        }
    }

    return best
}