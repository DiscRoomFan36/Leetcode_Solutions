package main

import "strings"

func decodeCiphertext(encodedText string, rows int) string {
    sb := strings.Builder{}

    cols := len(encodedText) / rows;

    for i := range cols {
        for j := range rows {
            col := i + j;
            // make sure we dont go out of bounds,
            // this could be calculated in the range.
            if col >= cols { break }

            index := j * cols + col;
            sb.WriteByte(encodedText[index]);
        }
    }

    result := sb.String()
    // make sure to remove all trailing " "
    return strings.TrimRight(result, " ");
}
