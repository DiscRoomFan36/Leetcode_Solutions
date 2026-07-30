package main

func minimumPushes(word string) int {
    distinct_letters := len(word);
    total_key_presses := 0;
    
    const NUM_KEYS = 8;
    for factor := 1; ; factor += 1 {
        if distinct_letters >= NUM_KEYS {
            total_key_presses += NUM_KEYS * factor;
            distinct_letters -= NUM_KEYS;
        } else {
            total_key_presses += distinct_letters * factor;
            return total_key_presses;
        }
    }
}
