package main

import (
	"slices"
	"strings"
)

func maxNumOfSubstrings(s string) []string {
    result := make([]string, 0);

    orginal_counts := get_character_count(s);
    poisoned_chars := [26]bool{};

    for index := 0; index < len(s); {
        c := s[index];
        total_amount_of_c := orginal_counts.count_of(c);

        // check if all of the same character is in a row
        count_of_c := 1;
        for i := index + 1; i < len(s); i += 1 {
            if s[i] != c { break; }
            count_of_c += 1;
        }
        if count_of_c == total_amount_of_c {
            // this means they are all in a row, so they are definently in the result
            Append(&result, s[index:index + total_amount_of_c]);

            poisoned_chars[c - 'a'] = true;

            // move start of next segment, and index
            index += total_amount_of_c;

        } else {
            // else its a character that is spred out. might still contain a valid sequnce, thats why segments exists.
            index += count_of_c;
        }
    }

    remove_poisoned_segments := func(segments []string) []string {
        for {
            // now i need to check the segments to see if they conatin any valid sequences
            running_counts := Bag{};
            for segment_index := 0; segment_index < len(segments); segment_index += 1 {
                segment := segments[segment_index];
                segment_counts := get_character_count(segment);

                for i := range segment_counts {
                    // if something is in both this segment, and a previous segment,
                    // its means there is a split, and this char cannot possibly be  
                    if segment_counts[i] > 0 && running_counts[i] > 0 {
                        poisoned_chars[i] = true;
                    }
                }
                running_counts = add_bags(running_counts, segment_counts);
            }

            // keep spliting on the posioned chars
            new_segments := make([]string, 0);

            poisoned_chars_array := make([]string, 0, 26);
            for i, is_poisoned := range poisoned_chars {
                if !is_poisoned { continue; }
                Append(&poisoned_chars_array, string(i + 'a'));
            }

            got_a_split := false;

            for _, segment := range segments {
                split := SplitAny(segment, poisoned_chars_array...);
                if slices.Compare(split, []string{segment}) != 0 {
                    got_a_split = true;
                }

                for _, subsegment := range split {
                    if len(subsegment) == 0 { continue; }
                    Append(&new_segments, subsegment);
                }
            }

            // replace old segments, time for another loop.
            if !got_a_split { break; }

            segments = new_segments;
        }

        return segments;
    }

    segments := remove_poisoned_segments([]string{s});

    for len(segments) > 0 {
        for _, segment := range segments {
            // find smallest valid
            //
            // segment is allways valid (or contains a smaller valid)
            smallest_valid := segment;
            // cacheing so we dont do n^2 loop.
            past_count := Bag{};

            for start := 0; start < len(segment); start += 1 {
                c := segment[start];
                if past_count.count_of(c) > 0 { continue; }
                past_count.add_rune(rune(c));

                running_counts := Bag{};
                for end := start; end < len(segment) && end < start + len(smallest_valid); end += 1 {
                    c := segment[end];
                    running_counts.add_rune(rune(c));

                    is_valid := true;
                    for i := range running_counts {
                        if running_counts[i] == 0 { continue; }
                        if running_counts[i] != orginal_counts[i] {
                            is_valid = false;
                            break;
                        }
                    }

                    if is_valid {
                        valid_segment := segment[start:end+1];
                        if len(valid_segment) < len(smallest_valid) {
                            smallest_valid = valid_segment;
                        }
                        break;
                    }
                }
            }

            Append(&result, smallest_valid);

            pre_poisoned := poisoned_chars;
            // all chars i just added are now poisoned. this could be faster, but eh.
            for i := range smallest_valid {
                poisoned_chars[smallest_valid[i]-'a'] = true;;
            }

            if pre_poisoned == poisoned_chars { panic("not possible"); }
        }

        // this operation is guarenteed to make the segments smaller,
        //
        // would assert here but that would be slow.
        segments = remove_poisoned_segments(segments);
    }

    return result;
}

func Append[T any](slice *[]T, items ...T) *T {
	*slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}

type Bag [26]int;

func (bag *Bag) add_rune(c rune) {
    bag[c - 'a'] += 1;
}

func (bag Bag) count_of(c byte) int {
    return bag[c - 'a'];
}

func get_character_count(s string) Bag {
    result := Bag{};
    for _, c := range s {
        result.add_rune(c);
    }
    return result;
}

func add_bags(a, b Bag) Bag {
    c := Bag{};
    for i := range c {
        c[i] = a[i] + b[i];
    }
    return c;
}

// yoinked from:
// https://github.com/projectdiscovery/utils/blob/v0.11.4/strings/stringsutil.go#L185
//
// also, now that i look at it, its got a couple bugs in it.
//
// SplitAny string by a list of separators
func SplitAny(s string, seps ...string) []string {
	sepsStr := strings.Join(seps, "")
	splitter := func(r rune) bool {
		return strings.ContainsRune(sepsStr, r)
	}
	return strings.FieldsFunc(s, splitter)
}
