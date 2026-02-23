package main

func hasAllCodes(s string, k int) bool {
    // there is probably some equation that gets the minimum anount of length required for this.
    if len(s) < k*2 { return false; }

    rolling_number := 0;

    // & rolling number with this.
    mask := (1 << k) - 1;

    seen_number := make([]bool, 1 << k);

    for i := range k-1 {
        rolling_number = (rolling_number << 1);
        if s[i] == '1' { rolling_number |= 1; }
    }

    for i := k-1; i < len(s); i++ {
        rolling_number = (rolling_number << 1) & mask;
        if s[i] == '1' { rolling_number |= 1; }
        seen_number[rolling_number] = true;
    }

    for _, seen := range seen_number {
        if !seen { return false; }
    }
    return true;
}