package main

func minAddToMakeValid(s string) int {
    result := 0;
    counter := 0;
    for _, c := range s {
        if c == '(' {
            counter += 1;
        } else {
            if c != ')' { panic("what"); }

            if counter == 0 {
                // need to add something before to balence this
                result += 1;
            } else {
                counter -= 1;
            }
        }
    }

    // add the counter because we need to close opened brackets
    return result + counter;
}
