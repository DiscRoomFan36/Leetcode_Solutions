package main

func maxActiveSectionsAfterTrade(s string) int {
    // remove outer 1's
    extra_active := 0;
    for len(s) > 0 && s[0] == '1' {
        s = s[1:];
        extra_active += 1;
    }
    for len(s) > 0 && s[len(s)-1] == '1' {
        s = s[:len(s)-1];
        extra_active += 1;
    }

    // there is nothing we can do with such a small s,
    if len(s) <= 2 { return extra_active; }

    // the only other way to make no trade good is if there are no extra 1's in the string
    is_a_one := false;
    for _, c := range s {
        if c == '1' { is_a_one = true; break; }
    }
    if !is_a_one { return extra_active; }


    // now we must find the optimal trade.
    //
    // step 1, find all possible trades.
    type Trade struct { start, end int; };
    trades := []Trade{};

    total_active := 0;
    for i := 0; i < len(s); i++ {
        if s[i] == '1' {
            total_active += 1;

            start := i;
            end := i+1
            for ; end < len(s); end++ {
                if s[end] == '0' { break; }
                total_active += 1;
            }
            if end == len(s) { panic("both sides are surrounded by zero's, this cannot happen."); }
            trades = append(trades, Trade{start, end});

            i = end-1;
        }
    }

    // step 2, find the best trade to make, 
    best_active := total_active;
    for i, trade := range trades {
        // this is the index of the 0 after the 1
        one_before := 0;
        if i > 0 { one_before = trades[i-1].end; }
        // this is the index of the 1 after the zeros
        one_after := len(s);
        if i < len(trades)-1 { one_after = trades[i+1].start; }

        // original number of 1's, to subtract.
        ones_in_a_row := trade.end - trade.start;
        // new number on 1's, to add
        new_run_of_ones := one_after - one_before;

        new_active := total_active - ones_in_a_row + new_run_of_ones;
        best_active = max(best_active, new_active);
    }

    // remember to add the extra we took away at the begining!
    return best_active + extra_active;
}
