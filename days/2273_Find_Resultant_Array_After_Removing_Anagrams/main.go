package main

func removeAnagrams(words []string) []string {
    // bucket contains the count of letters in the word.
    // anagrams conatin the same amount of letter frequency.
    type Bucket ['z' - 'a' + 1]int

    // dont be weird about it.
    safe_words := []string{}
    last_bucket := Bucket{}

    for i, word := range words {
        // ha ha.
        if i == 0 && word == "" { safe_words = append(safe_words, word); continue }

        // construct the bucket 
        bucket := Bucket{}
        // will panic if the word conatins letters outside of [a -> z]
        for _, c := range word { bucket[c - 'a'] += 1 }

        if bucket != last_bucket {
            last_bucket = bucket
            safe_words = append(safe_words, word)
        }
    }

    return safe_words
}