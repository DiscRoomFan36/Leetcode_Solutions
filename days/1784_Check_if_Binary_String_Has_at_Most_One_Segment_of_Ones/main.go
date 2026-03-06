package main

import "strings"

func checkOnesSegment(s string) bool {
    if s[0] == '0' { panic("no leading zeros"); }

    first_zero := strings.IndexByte(s, '0');
    if first_zero == -1 { return true; }

    other_segment_of_one := strings.IndexByte(s[first_zero+1:], '1');
    return other_segment_of_one == -1;
}
