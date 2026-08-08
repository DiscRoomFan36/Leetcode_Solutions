package main

func validSequence(word1 string, word2 string) []int {
    last := make([]int, len(word2));
    for i := range last { last[i] = -1; }

    { // populate last array
        j := len(word2) - 1;
        for i := len(word1)-1; i >= 0 && j >= 0; i -= 1 {
            if word1[i] == word2[j] {
                last[j] = i;
                j -= 1;
            }
        }
    }

    result := []int{};
    have_skip := true;
    j := 0;
    for i := 0; i < len(word1); i += 1 {
        if j == len(word2) { break; }
        letter_equal := word1[i] == word2[j];
        // i think i understand this line after enough stareing.
        skip_avaliable := have_skip && (j == len(word2)-1 || i < last[j+1]);
        if letter_equal || skip_avaliable {
            // check if we just used a skip.
            if word1[i] != word2[j] { have_skip = false; }
            result = append(result, i);
            j += 1;
        }
    }
    if j == len(word2) { return result }
    return []int{};
}
