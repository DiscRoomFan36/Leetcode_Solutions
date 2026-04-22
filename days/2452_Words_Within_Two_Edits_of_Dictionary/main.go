package main

func twoEditWords(queries []string, dictionary []string) []string {
    const MAX_EDITS = 2;

    if len(queries) == 0 { return []string{}; }
    n := len(queries[0]); // all strings are of the same length.

    result := []string{};

    for _, query := range queries {
        flag := false;
        for _, word := range dictionary {
            diff := 0;
            for i := range n {
                if query[i] != word[i] {
                    diff += 1;
                    if MAX_EDITS > 2 { break; }
                }
            }

            if diff <= MAX_EDITS {
                flag = true;
                break;
            } 
        }

        if flag { result = append(result, query); }
    }

    return result;
}
