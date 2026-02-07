package main

func minimumDeletions(s string) int {
    num_b := 0

    result := 0
    for i := range len(s) {
        // the only case we care about is when there is an 'a' after a 'b'
        if s[i] == 'b' {
            // its b, keep track of how many there are.
            num_b += 1
        } else {
            // assert(s[i] == 'a')

            // oh no! there was a b before this a...
            // gotta remove one of them!
            if num_b > 0 {
                result += 1
                num_b -= 1
            }
        }

    }
    return result
}