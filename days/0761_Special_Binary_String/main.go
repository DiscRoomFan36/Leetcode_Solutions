package main

func makeLargestSpecial(s string) string {
    //
    // man, go fuck yourself.
    //
    // please explain were in the insrtuctions that we
    // are allowed to swap more than one piece? im waiting.
    //
    // also the leetcode solution is dumb.
    //
    best_string := s;
    old_best := ""
    for best_string != old_best {
        old_best = best_string;
        s = best_string;

        // dumb O(n^3) algo, but its optmized, so its probably faster than whatever the real solution is.
        for k := 0; k < len(s); k++ {
            found_bigger_this_round := false;

            // we are going check if this is special
            // along the way, will be faster.
            str1_num_ones  := 0;
            str1_num_zeros := 0;

            // skip by 2 at a time, special strings are allways even.
            for i := k+2; i < len(s); i += 2 {

                if s[i-2] == '1' { str1_num_ones  += 1;
                } else           { str1_num_zeros += 1; }
                if str1_num_ones < str1_num_zeros { break; }

                if s[i-1] == '1' { str1_num_ones  += 1;
                } else           { str1_num_zeros += 1; }
                if str1_num_ones < str1_num_zeros { break; }

                if str1_num_ones != str1_num_zeros { continue; }

                str1 := s[k:i]

                // we are going check if this is special
                // along the way, will be faster.
                str2_num_ones  := 0;
                str2_num_zeros := 0;

                // skip by 2 at a time, special strings are allways even.
                for j := i+1; j < len(s); j += 2 {

                    if s[j-1] == '1' { str2_num_ones  += 1;
                    } else           { str2_num_zeros += 1; }
                    if str2_num_ones < str2_num_zeros { break; }

                    if s[j  ] == '1' { str2_num_ones  += 1;
                    } else           { str2_num_zeros += 1; }
                    if str2_num_ones < str2_num_zeros { break; }
                    
                    if str2_num_ones != str2_num_zeros { continue; }

                    str2 := s[i:j+1]

                    new_string := s[:k] + str2 + str1 + s[j+1:];

                    if best_string < new_string {
                        best_string = new_string;
                        found_bigger_this_round = true;
                    }

                }
            }

            // this works becase we got the first instance, and no further numbers can make anything bigger.
            if found_bigger_this_round { break; }
        }

    }

    return best_string;
}
