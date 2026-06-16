package main

func processStr(s string) string {
    result := make([]byte, 0);
    for _, c := range s {
        if 'a' <= c && c <= 'z' {
            // just add character
            Append(&result, byte(c));
        } else if c == '*' {
            // remove last
            if len(result) > 0 { result = result[:len(result)-1]; }
        } else if c == '#' {
            // dublicate it
            if len(result) > 0 { Append(&result, result...); }
        } else if c == '%' {
            // reverse
            for i := range len(result)/2 {
                j := len(result)-1 - i;
                result[i], result[j] = result[j], result[i];
            }
        } else {
            panic("unreachable");
        }
    }
    // this is a copy in golang i think.
    return string(result);
}

func Append[T any](slice *[]T, items ...T) *T {
	*slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}
